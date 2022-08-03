package main

import (
	"bytes"
	"encoding/json"
	_"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
)

func getDirectPairReserves(tokenA common.Address, tokenB common.Address, symbolA string, symbolB string) map[string]map[string]string {

	data := `{"query":"{block{`

	for key, value := range dexList {
		data += symbolA+symbolB+key+`:account(address:\"`+crate2(tokenA, tokenB, value).Hex()+`\"){storage(slot:\"0x0000000000000000000000000000000000000000000000000000000000000008\")address}`
		data += symbolA+symbolB+key+`:account(address:\"`+crate2(tokenB, tokenA, value).Hex()+`\"){storage(slot:\"0x0000000000000000000000000000000000000000000000000000000000000008\")address}`
	}
	data += `}}"}`

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