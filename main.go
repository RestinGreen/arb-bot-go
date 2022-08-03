package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
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

	user := os.Getenv("DB_USER")
	pass := os.Getenv("PASS")
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	name := os.Getenv("NAME")
	initDB(user, pass, host, port, name)

}


func main() {
	ctx := context.Background()
	txs := make(chan *types.Transaction)

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
	_, err = subscriber.SubscribePendingTransactions(ctx, txs)
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

	for tx := range txs {
		
		for dex, data := range dexList {
			if tx.To() != nil && tx.To().Hash() == data.Router.Hash() {
				log := ""
				buildLog(&log, dex+"swap")
				method, err := uniV2Abi.MethodById(tx.Data()[0:4])
				if err != nil {
					fmt.Println(err)
					break
				}
				switch method.Name {
					case "swapExactTokensForTokens":
						swapWithSimpleFeeType1(tx, method, log, data)
					case "swapTokensForExactTokens":
						swapWithSimpleFeeType1(tx, method, log, data)
					case "swapTokensForExactETH":
						swapWithSimpleFeeType1(tx, method, log, data)
					case "swapExactTokensForETH":
						swapWithSimpleFeeType1(tx, method, log, data)
					case "swapExactETHForTokens":
						swapWithSimpleFeeType2(tx, method, log, data)
					case "swapETHForExactTokens":
						swapWithSimpleFeeType2(tx, method, log, data)
					default:
					
				}
				break
			}
		}

	}

}
