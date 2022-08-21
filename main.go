package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	
)

func init() {

	initBot()

}

func cleanup() {
	DB.Close()
}

func main() {

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		cleanup()
		fmt.Println("\nEXITING")
		os.Exit(1)
	}()

	ctx := context.Background()
	txs := make(chan *types.Transaction)

	

	subscriber := gethclient.New(baseClient)
	
	_, err := subscriber.SubscribePendingTransactions(ctx, txs)
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
				log := "----------------------------------------------------------------------------------------\n"
				buildLog(&log, dex + " swap")
				noticed := time.Now()
				if len(tx.Data()) < 4 {
					break
				}
				method, err := uniV2Abi.MethodById(tx.Data()[0:4])
				if err != nil {
					fmt.Println(err)
					break
				}
				switch method.Name {
				case "swapExactTokensForTokens":
					prefetchTxData(tx, method, log, data, 1, noticed)
				case "swapTokensForExactTokens":
					prefetchTxData(tx, method, log, data, 2, noticed)
				case "swapTokensForExactETH":
					prefetchTxData(tx, method, log, data, 3, noticed)
				case "swapExactTokensForETH":
					prefetchTxData(tx, method, log, data, 4, noticed)
				case "swapExactETHForTokens":
					prefetchTxData(tx, method, log, data, 5, noticed)
				case "swapETHForExactTokens":
					prefetchTxData(tx, method, log, data, 6, noticed)
				default:

				}
				break
			}
		}
	}

}
