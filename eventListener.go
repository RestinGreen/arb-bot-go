package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"main/ReservesFetcher"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

var finalReserves = make(map[common.Address]ReservesFetcher.ReservesFetcherTokenReserve)

func getAllDirectPairReserves(tokens []TokenPairUsage) []ReservesFetcher.ReservesFetcherTokenReserve {
	pairs := make([]common.Address, 0)
	for _, t := range tokens {
		for _, data := range dexList {
			crate2 := crate2(common.HexToAddress(t.AddressA), common.HexToAddress(t.AddressB), data)
			pairs = append(pairs, crate2)
		}
	}

	callOpts := bind.CallOpts{
		Pending: false,
		Context: nil,
	}
	start := time.Now()
	reserves, err := fetcherContract.Fetch(&callOpts, pairs)
	fmt.Println(time.Since(start))
	if err != nil {
		panic(err)
	}
	fmt.Println(len(reserves))

	return reserves

}

func prepareListening() {

	tokens := getTokenUsage()

	reserves := getAllDirectPairReserves(tokens)

	finalWatchList := make([]common.Address, 0)
	for _, elem := range reserves {
		if elem.ReserveA.Cmp(ZERO) != 0 && elem.ReserveB.Cmp(ZERO) != 0 {
			finalReserves[elem.Pair] = elem
			finalWatchList = append(finalWatchList, elem.Pair)
		}
	}

	query := ethereum.FilterQuery{
		Addresses: finalWatchList,
		Topics:    [][]common.Hash{{common.HexToHash("0x1c411e9a96e071241c2f21f7726b17ae89e3cab4c78be50e062b03a9fffbbad1")}},
	}

	logs := make(chan types.Log)

	subscribe, err := ethClient.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	go startListening(subscribe, logs)
}

func startListening(subscribe ethereum.Subscription, logs chan types.Log) {
	for {
		select {
		case err := <-subscribe.Err():
			log.Fatal(err)
		case log := <-logs:
			r0 := new(big.Int).SetBytes(log.Data[0:32])
			r1 := new(big.Int).SetBytes(log.Data[32:64])
			if entry, ok := finalReserves[log.Address]; ok {
				if entry.BlockNumber.Uint64() < log.BlockNumber {
					entry.ReserveA = r0
					entry.ReserveB = r1
					entry.TxIndex = ZERO
					finalReserves[log.Address] = entry
				} else if entry.BlockNumber.Uint64() == log.BlockNumber {
					if uint(entry.TxIndex.Uint64()) <= log.TxIndex {
						entry.ReserveA = r0
						entry.ReserveB = r1
						entry.TxIndex = big.NewInt(int64(log.TxIndex))
						finalReserves[log.Address] = entry
					}
				}
			}
			// fmt.Println(finalReserves[log.Address].Pair, finalReserves[log.Address].BlockNumber, finalReserves[log.Address].TxIndex)
		}
	}

}