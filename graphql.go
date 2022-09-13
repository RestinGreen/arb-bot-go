package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"math/big"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

func getTriggerTxReserves(path []common.Address, dexData DexData) map[string]map[string]string {

	data := `{"query":"{block{`
	for i := 0; i < len(path)-1; i++ {
		tokenA := path[i].String()
		tokenB := path[i+1].String()
		data += tokenA[1:] + tokenB[1:] + `:account(address:\"`+crate2(path[i], path[i+1], dexData).Hex()+`\"){storage(slot:\"0x0000000000000000000000000000000000000000000000000000000000000008\")address}`
	}
	data += `}}"}`
	return makeGQLCall(data)
}

func getDirectPairReserves(tokenA common.Address, tokenB common.Address, symbolA string, symbolB string) map[string]map[string]string {

	gqlQueryString, ok := gqlQueryStrings[tokenA.String()+tokenB.String()]
	if !ok {
		// fmt.Println("not found in fetched query strings")

		data := `{"query":"{block{`
		allPairs := make(map[string]string)
		goodPairs := make(map[string]string)
		for dexName, dexData := range dexList {
			newKey := strings.ToUpper(symbolA + "_" + symbolB + "_" + dexName)
			line := newKey + `:account(address:\"`+crate2(tokenA, tokenB, dexData).Hex()+`\"){storage(slot:\"0x0000000000000000000000000000000000000000000000000000000000000008\")address}`
			data += line
			allPairs[newKey] = line
		}
		data += `}}"}`
		result := makeGQLCall(data)
		for pair, pairData := range(result) {
			s := new(big.Int)
			s.SetString(pairData["storage"][2:], 16)
			if s.Cmp(ZERO) != 0 {
				goodPairs[pair] = allPairs[pair]
			}
		}
		goodGQLQueryString := `{"query":"{block{`
		for _, sqlString := range(goodPairs) {
			goodGQLQueryString += sqlString
		}
		goodGQLQueryString += `}}"}`

		insertGQLQueryString(tokenA.String()+tokenB.String(), goodGQLQueryString)
		return makeGQLCall(goodGQLQueryString)


	} else {
		// fmt.Println("found in fetched...")
		return makeGQLCall(gqlQueryString.queryString)
	}
}


func makeGQLCall(data string) map[string]map[string]string{
	bytesData := []byte(data)
	
	body := bytes.NewBuffer(bytesData)
	
	req, err := http.NewRequest("POST", "http://localhost:8545/graphql", body)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	
	bdy, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	jsonStr := string(bdy)
	
	var jsonMap map[string]map[string]map[string]map[string]string
	
	json.Unmarshal([]byte(jsonStr), &jsonMap)
	return jsonMap["data"]["block"]
}