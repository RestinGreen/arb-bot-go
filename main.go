package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/joho/godotenv"
)



func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	wsEndpoint = os.Getenv("NODE_WS_ENDPOINT")
	ipcEndpoint = os.Getenv("NODE_IPC_ENDPOINT")
	dexJson = os.Getenv("DEX_JSON")
	uniV2Json = os.Getenv("UNIV2_JSON")

	readDexFromFile(dexJson)
	readUniV2Router02(uniV2Json)


}

func graphqlTest() {
	fmt.Println("-----------------")

	// Generated by curl-to-Go: https://mholt.github.io/curl-to-go

	// curl -X POST \
	// -H "Content-Type: application/json" \
	// -d '{"query": "{block{number}}"}' \
	// http://localhost:8545/graphql

	data := []byte(`{"query":"{block{WETHUSDCSUSHISWAP:account(address:\"0x20f8a5947367e3b42fa2c2a5973d3780c505cd58\"){storage(slot:\"0x0000000000000000000000000000000000000000000000000000000000000008\")address}}}"}`)
	//  data := []byte(`{"query":"{block{WETHUSDCSUSHISWAP:account(address:\"0x20f8a5947367e3b42fa2c2a5973d3780c505cd58\"){storage(slot:\"0x0000000000000000000000000000000000000000000000000000000000000008\")address}}}"}`)

	body := bytes.NewBuffer(data)

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
	fmt.Println(jsonMap["data"]["block"]["WETHUSDCSUSHISWAP"]["address"])

}

func main() {
	ctx := context.Background()
	txHash := make(chan common.Hash)

	baseClient, err := rpc.Dial(wsEndpoint)
	if err != nil {
		panic(err)
	}
	ethClient, err = ethclient.Dial(ipcEndpoint)
	if err != nil {
		panic(err)
	}
	
	subscriber := gethclient.New(baseClient)
	if err != nil {
		panic(err)
	}
	_, err = subscriber.SubscribePendingTransactions(ctx, txHash)
	if err != nil {
		panic(err)
	}
	defer func() {
		baseClient.Close()
		ethClient.Close()
	}()

	uniV2Abi, err := abi.JSON(strings.NewReader(uniV2ABIString))
	if err != nil {
		panic(err)
	}

	for txHash := range txHash {
		tx, _, err := ethClient.TransactionByHash(ctx, txHash)
		if err != nil {
			fmt.Println(err)
			continue
		}

		for dex, data := range dexList {
			if tx.To() != nil && tx.To().Hash() == data.Router.Hash() {
				log := ""
				buildLog(&log, dex+"swap")
				method, err := uniV2Abi.MethodById(tx.Data()[0:4])
				if err != nil {
					panic(err)
				}
				
				switch method.Name {
				case "swapExactTokensForTokens":
					swapWithSimpleFeeType1(tx, method, log)
				case "swapTokensForExactTokens":
					swapWithSimpleFeeType1(tx, method, log)
				case "swapTokensForExactETH":
					swapWithSimpleFeeType1(tx, method, log)
				case "swapExactTokensForETH":
					swapWithSimpleFeeType1(tx, method, log)
				case "swapExactETHForTokens":
					swapWithSimpleFeeType2(tx, method, log)
				case "swapETHForExactTokens":
					swapWithSimpleFeeType2(tx, method, log)
				default:
					
				}
				break
			}
		}

	}

}
