package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	// "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/mitchellh/mapstructure"
)

type UniV2Outputs struct {
	Amounts []*big.Int
}

type UniV2Inputs struct {
	AmountIn     *big.Int
	AmountInMax  *big.Int
	AmountOut    *big.Int
	AmountOutMin *big.Int
	Path         []common.Address
	To           common.Address
	Deadline     *big.Int
}

func prefetchTxData(tx *types.Transaction, method *abi.Method, log string, dexData DexData, txType int, noticed time.Time) {
	buildLog(&log, method.Name)

	var input UniV2Inputs
	inputMap := make(map[string]interface{}, 0)
	start := time.Now()
	err := method.Inputs.UnpackIntoMap(inputMap, tx.Data()[4:])
	if err != nil {
		panic(err)
	}
	fmt.Println("unpack time ", time.Since(start))
	start = time.Now()
	err = mapstructure.Decode(inputMap, &input)
	fmt.Println("decode time ", time.Since(start))
	if err != nil {
		panic(err)
	}

	// sender, err := types.Sender(types.NewEIP155Signer(tx.ChainId()), tx)
	// if err != nil {
	// 	sender, err = types.Sender(types.NewLondonSigner(tx.ChainId()), tx)
	// }
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// callMsg := ethereum.CallMsg{
	// 	To:   tx.To(),
	// 	Data: tx.Data(),
	// 	From: sender,
	// 	Value: tx.Value(),
	// }

	start = time.Now()
	queryResult := getTriggerTxReserves(input.Path, dexData)
	decodedReserves := make([]PairReserve, len(input.Path))
	for i := 0; i < len(queryResult); i++ {
		key := input.Path[i].String()[1:] + input.Path[i+1].String()[1:]
		a, b, ok := decodeStorageSlot(queryResult[key]["storage"], input.Path[i], input.Path[i+1])
		if !ok {
			return
		}
		decodedReserves[i] = PairReserve{tokenA: a, tokenB: b}

	}
	var amounts []*big.Int
	var ok bool
	switch txType {
	case 1:
		amounts, ok = simulateSwapExactTokensForTokens(input.AmountIn, input.Path, decodedReserves)
	case 2:
		amounts, ok = simulateSwapTokensForExactTokens(input.AmountOut, input.Path, decodedReserves)
	case 3:
		amounts, ok = simulateSwapTokensForExactETH(input.AmountOut, input.Path, decodedReserves)
	case 4:
		amounts, ok = simulateSwapExactTokensForETH(input.AmountIn, input.Path, decodedReserves)
	case 5:
		amounts, ok = simulateSwapExactETHForTokens(tx.Value(), input.Path, decodedReserves)
	case 6:
		amounts, ok = simulateSwapETHForExactTokens(input.AmountOut, input.Path, decodedReserves)
	}
	if !ok {
		return
	}
	
	buildLog(&log, time.Since(start).String())
	fmt.Println("main reserver time ", time.Since(start))

	// start := time.Now()
	//simulation, err := ethClient.CallContract(context.Background(), callMsg, nil)
	//buildLog(&log, "simulation time "+time.Since(start).String())
	// if err != nil {
	// 	fmt.Println(tx.Hash())
	// 	fmt.Println(err)
	// 	return
	// }
	// var output UniV2Outputs
	// outputMap := make(map[string]interface{}, 0)
	// err = method.Outputs.UnpackIntoMap(outputMap, simulation)
	// if err != nil {
	// 	panic(err)
	// }

	// err = mapstructure.Decode(outputMap, &output)
	// if err != nil {
	// 	panic(err)
	// }
	for i := 0; i < len(input.Path)-1; i++ {

		go checkArbitrage(input.Path[i], input.Path[i+1], amounts[i], amounts[i+1], tx, dexData.Name, log, noticed)
	}

}

func checkArbitrage(tokenA common.Address, tokenB common.Address, tokenASoldAmount *big.Int, tokenBBoughtAmount *big.Int, tx *types.Transaction, skipDex string, log string, noticed time.Time) {

	fmt.Println(time.Now())
	start := time.Now()
	A, foundA := getTokenDataByAddress(tokenA.Hex())
	B, foundB := getTokenDataByAddress(tokenB.Hex())
	fmt.Println("perload ", time.Since(start))
	buildLog(&log, "big trade "+tx.Hash().Hex())
	buildLog(&log, "sold\t"+A.Symbol+"\t"+tokenASoldAmount.String())
	buildLog(&log, "bought\t"+B.Symbol+"\t"+tokenBBoughtAmount.String())

	if !foundA || !foundB {
		return
	}

	maxProfit := ZERO
	profit := ZERO
	var arbParams ArbParams

	pair1Key := strings.ToUpper(A.Symbol + "_" + B.Symbol + "_" + skipDex)

	gqlTime := time.Now()
	reserves := getDirectPairReserves(tokenA, tokenB, A.Symbol, B.Symbol)
	buildLog(&log, "graphql time: "+time.Since(gqlTime).String())
	reserveA1, reserveB1, ok := decodeStorageSlot(reserves[pair1Key]["storage"], tokenA, tokenB)
	if !ok {
		fmt.Println(skipDex, "needs investigation, maybe not using crate2 for token pair cration")
		fmt.Println(reserves[pair1Key]["address"])
		return
	}
	pair1 := reserves[pair1Key]["address"]
	simulatedReserveA := new(big.Int).Add(reserveA1, tokenASoldAmount)
	simulatedReserveB := new(big.Int).Sub(reserveB1, tokenBBoughtAmount)
	if simulatedReserveB.Cmp(ZERO) < 0 {
		fmt.Println(A.Address)
		fmt.Println(B.Address)
		fmt.Println("pool ", pair1)
		fmt.Println("reserve a1 ", reserveA1)
		fmt.Println("reserve b1 ", reserveB1)
		fmt.Println("sold ", tokenASoldAmount, " ", A.Symbol)
		fmt.Println("bought ", tokenBBoughtAmount, " ", B.Symbol)
		fmt.Println("dex: ", skipDex)
		fmt.Println(reserves)
		return
	}

	for pair := range reserves {
		split := strings.Split(pair, "_")
		dex := split[2]
		if len(split) > 3 {
			dex += "_" + split[3]
		}
		if dex != skipDex {
			pair2Key := strings.ToUpper(A.Symbol + "_" + B.Symbol + "_" + dex)
			storageSlot, ok := reserves[pair2Key]["storage"]
			if !ok {
				fmt.Println("ss not fount")
			}
			tmp := new(big.Int)
			tmp.SetString(storageSlot[2:], 16)

			//skip if token pair doesnt exist for given DEX
			if tmp.Cmp(ZERO) == 0 {
				continue
			}
			reserveA2, reserveB2, ok2 := decodeStorageSlot(storageSlot, tokenA, tokenB)
			if !ok2 {
				fmt.Println(dex)
				continue
			}
			pair2 := reserves[pair2Key]["address"]
			optimalInput := calculateOptimalInput(simulatedReserveB, simulatedReserveA, reserveA2, reserveB2)

			if optimalInput.Cmp(ZERO) > 0 {
				profit = calculateProfit(optimalInput, simulatedReserveB, simulatedReserveA, reserveA2, reserveB2)
				if profit.Cmp(ZERO) == 0 {
					continue
				}
				buildLog(&log, "profit "+profit.String()+" "+B.Symbol+" "+dex)
				if profit.Cmp(new(big.Int).Div(new(big.Int).Mul(optimalInput, big.NewInt(3)), b1000)) > 0 {
					maxProfit = profit
					flashAmount, ok := getAmountOut(simulatedReserveB, simulatedReserveA, optimalInput)
					if !ok {
						return
					}
					arbParams = ArbParams{
						flashToken:   tokenA,
						flashAmount:  flashAmount,
						paybackToken: tokenB,
						path:         []common.Address{common.HexToAddress(pair1), common.HexToAddress(pair2)},
						dex:          dex,
					}
				}
			}

		}
	}

	minimum := new(big.Int)
	minimum.SetString(B.Minimum, 10)
	buildLog(&log, "minimum profit: "+B.Minimum+" "+B.Symbol)
	if maxProfit.Cmp(minimum) > 0 {
		cost := new(big.Int)

		if tx.Type() == 0 {
			cost.Mul(tx.GasPrice(), bn150k).Mul(cost, big.NewInt(4))
		} else if tx.Type() == 2 {
			cost.Mul(tx.GasFeeCap(), bn150k).Mul(cost, big.NewInt(4))
		}

		cost.Mul(cost, TEN.Exp(TEN, big.NewInt(18-B.decimals), nil))
		if cost.Cmp(maxProfit) < 0 {
			go executeArbitrage(arbParams, tx, fromAddress0, privateKey0, noticed)
			go executeArbitrage(arbParams, tx, fromAddress1, privateKey1, noticed)
			go executeArbitrage(arbParams, tx, fromAddress2, privateKey2, noticed)
			go executeArbitrage(arbParams, tx, fromAddress3, privateKey3, noticed)
			go executeArbitrage(arbParams, tx, fromAddress4, privateKey4, noticed)
			buildLog(&log, "input: "+arbParams.flashAmount.String())
			buildLog(&log, "pair1: "+arbParams.path[0].String())
			buildLog(&log, "pair2: "+arbParams.path[1].String())
			buildLog(&log, "\033[32mprofit: "+maxProfit.String()+" "+B.Symbol+" @ "+arbParams.dex+"\033[0m")
			fmt.Println(log)
		}
	}

}

func executeArbitrage(params ArbParams, tx *types.Transaction, sender common.Address, privateKey *ecdsa.PrivateKey, noticed time.Time) {
	nonce, err := ethClient.PendingNonceAt(context.Background(), sender)
	if err != nil {
		log.Fatal(err)
		return
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal(err)
		return
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasLimit = uint64(500_000)
	if tx.Type() == 0 {
		auth.GasPrice = tx.GasPrice()
	} else {
		auth.GasFeeCap = tx.GasFeeCap()
		auth.GasTipCap = tx.GasTipCap()
	}

	transaction, err := contract.PrintMoney(auth, params.flashAmount, params.flashToken, params.paybackToken, params.path)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("transaction ", transaction.Hash().Hex(), "\n", "run time: ", time.Since(noticed))


}
