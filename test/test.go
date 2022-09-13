package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type UniV2Inputs struct {
	AmountIn     *big.Int
	AmountInMax  *big.Int
	AmountOut    *big.Int
	AmountOutMin *big.Int
	Path         []common.Address
	To           common.Address
	Deadline     *big.Int
}

func sortAddtess(tokenA common.Address, tokenB common.Address) (common.Address, common.Address) {

	if tokenA.Hash().Big().Cmp(tokenB.Hash().Big()) < 0 {
		return tokenA, tokenB
	} else {
		return tokenB, tokenA
	}
}

func crate2(tokenA common.Address, tokenB common.Address, data DexData) common.Address {

	tA, tB := sortAddtess(tokenA, tokenB)
	ff := []byte{255}
	ff = append(ff, data.Factory.Bytes()...)
	ff = append(ff, crypto.Keccak256(append(tA.Bytes(), tB.Bytes()...))...)
	ff = append(ff, common.Hex2Bytes(data.Salt)...)

	return common.BytesToAddress(crypto.Keccak256(ff))
}

type DexType int

const (
	UniV2 DexType = iota
	UniV3
	Curve
)

type DexData struct {
	Name    string
	Router  common.Address
	Factory common.Address
	Salt    string
	DexType DexType
}

var dexList = map[string]DexData{}

func readDexFromFile(jsonFileName string) {

	type TempDexData struct {
		Name    string `json:"name"`
		Router  string `json:"router"`
		Factory string `json:"factory"`
		Salt    string `json:"salt"`
		DexType string `json:"dexType"`
	}

	jsonFile, err := os.Open(jsonFileName)
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		panic(err)
	}
	var temp map[string]TempDexData
	json.Unmarshal([]byte(byteValue), &temp)

	for key, value := range temp {
		var dexType DexType
		if value.DexType == "UniV2" {
			dexType = UniV2
		} else if value.DexType == "UniV3" {
			dexType = UniV3
		} else {
			dexType = Curve
		}
		dexList[key] = DexData{
			Name:    value.Name,
			Router:  common.HexToAddress(value.Router),
			Factory: common.HexToAddress(value.Factory),
			Salt:    value.Salt,
			DexType: dexType,
		}
	}
}

func init() {
	readDexFromFile("dex.json")
}

func graphql() {

	start := time.Now()
	data := []byte(`{"query":"{block{WETHUSDCSUSHISWAP:account(address:\"0x20f8a5947367e3b42fa2c2a5973d3780c505cd58\"){storage(slot:\"0x0000000000000000000000000000000000000000000000000000000000000008\")address}}}"}`)

	body := bytes.NewBuffer(data)
	// crate2(usdc, weth, dexList["SUSHISWAP"])

	req, err := http.NewRequest("POST", "http://localhost:8545/graphql", body)
	if err != nil {
		// handle err
		fmt.Println("masodik")
		// fmt.Println(err)

	}
	req.Header.Set("Content-Type", "application/json")
	// i := 1
	// for i < 100 {
	// 	start := time.Now()
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
		fmt.Println("harmadik")
		// fmt.Println(err)

	}
	// 	fmt.Println(time.Since(start))
	// 	json.NewDecoder(resp.Body).Decode(&decoded)
	// }
	defer resp.Body.Close()

	bdy, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("negyedik")
	}
	jsonStr := string(bdy)
	// fmt.Println(jsonStr)

	var jsonMap map[string]map[string]map[string]map[string]string

	json.Unmarshal([]byte(jsonStr), &jsonMap)
	fmt.Println(time.Since(start))
	fmt.Println(jsonMap["data"])
}

func gql() {
	start := time.Now()
	data := `{"query":"{block{`
	data += `account(address:\"0xA0d6567bDaa90b996dACfe3140F16A45B9e66968\"){storage(slot:\"0x0000000000000000000000000000000000000000000000000000000000000008\")address}`
	data += `}}"}`
	bytesData := []byte(data)

	body := bytes.NewBuffer(bytesData)

	req, err := http.NewRequest("POST", "http://localhost:8545/graphql", body)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	fmt.Println(time.Since(start))
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
	fmt.Println(jsonMap["data"]["block"])
}

var uniV2ABIString string

func readUniV2Router02(fileName string) {
	fileBytes, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	uniV2ABIString = string(fileBytes)
}

func t(x *uint64) {
	*x = *x+1
}
func main() {

	var x uint64

	x = 1
	t(&x)
	fmt.Println(x)
}
