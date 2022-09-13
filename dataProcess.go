package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"time"

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
	// buildLog(&log, method.Name)

	var input UniV2Inputs
	inputMap := make(map[string]interface{}, 0)
	// start := time.Now()
	err := method.Inputs.UnpackIntoMap(inputMap, tx.Data()[4:])
	if err != nil {
		panic(err)
	}
	// fmt.Println("unpack time ", time.Since(start))
	// start = time.Now()
	err = mapstructure.Decode(inputMap, &input)
	// fmt.Println("decode time ", time.Since(start))
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

	// start = time.Now()
	// queryResult := getTriggerTxReserves(input.Path, dexData)
	// decodedReserves := make([]PairReserve, len(input.Path)-1)
	inputReserves := make([]PairReserve, len(input.Path)-1)
	for i := 0; i < len(input.Path)-1; i++ {
		crate2Address := crate2(input.Path[i], input.Path[i+1], dexData)
		pair, ok := finalReserves[crate2Address]
		if !ok {
			// fmt.Println("pair not used enough to make trades with it", crate2Address)
			return
		}
		if input.Path[i].Hash().Big().Cmp(input.Path[i+1].Hash().Big()) < 0 {
			inputReserves[i] = PairReserve{tokenA: pair.ReserveA, tokenB: pair.ReserveB}
		} else {
			inputReserves[i] = PairReserve{tokenA: pair.ReserveB, tokenB: pair.ReserveA}
		}
		// fmt.Println("pair ", pair.Pair)

	}
	// for i := 0; i < len(queryResult); i++ {
	// 	key := input.Path[i].String()[1:] + input.Path[i+1].String()[1:]
	// 	a, b, ok := decodeStorageSlot(queryResult[key]["storage"], input.Path[i], input.Path[i+1])
	// 	if !ok {
	// 		return
	// 	}
	// 	decodedReserves[i] = PairReserve{tokenA: a, tokenB: b}

	// }
	var amounts []*big.Int
	var ok bool
	switch txType {
	case 1:
		amounts, ok = simulateSwapExactTokensForTokens(input.AmountIn, input.AmountOutMin, input.Path, inputReserves)
	case 2:
		amounts, ok = simulateSwapTokensForExactTokens(input.AmountOut, input.AmountInMax, input.Path, inputReserves)
	case 3:
		amounts, ok = simulateSwapTokensForExactETH(input.AmountOut, input.AmountInMax, input.Path, inputReserves)
	case 4:
		amounts, ok = simulateSwapExactTokensForETH(input.AmountIn, input.AmountOutMin, input.Path, inputReserves)
	case 5:
		amounts, ok = simulateSwapExactETHForTokens(tx.Value(), input.AmountOutMin, input.Path, inputReserves)
	case 6:
		amounts, ok = simulateSwapETHForExactTokens(input.AmountOut, input.Path, inputReserves)
	}
	if !ok {
		return
	}

	// fmt.Println("gql")
	// for _, x := range decodedReserves {
	// 	fmt.Println(x.tokenA.String(), x.tokenB.String())
	// }
	// fmt.Println("preload")
	// for _, x := range inputReserves {
	// 	fmt.Println(x.tokenA.String(), x.tokenB.String())
	// }

	// fmt.Println("path ", input.Path)
	// fmt.Println("amounts ", amounts)

	// buildLog(&log, time.Since(start).String())
	// fmt.Println("main reserver time ", time.Since(start))

	// start := time.Now()
	// simulation, err := ethClient.CallContract(context.Background(), callMsg, nil)
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
	// fmt.Println("simulat ", output)
	for i := 0; i < len(input.Path)-1; i++ {

		go checkArbitrage(input.Path[i], input.Path[i+1], amounts[i], amounts[i+1], tx, dexData, log, noticed)
	}

}

func checkArbitrage(tokenA common.Address, tokenB common.Address, tokenASoldAmount *big.Int, tokenBBoughtAmount *big.Int, tx *types.Transaction, skipDex DexData, log string, noticed time.Time) {

	// fmt.Println(time.Now())
	// start := time.Now()
	A, foundA := getTokenDataByAddress(tokenA.Hex())
	B, foundB := getTokenDataByAddress(tokenB.Hex())

	// fmt.Println("preload ", time.Since(start))#
	// buildLog(&log, "big trade "+tx.Hash().Hex())
	// buildLog(&log, "sold\t"+A.Symbol+"\t"+tokenASoldAmount.String())
	// buildLog(&log, "bought\t"+B.Symbol+"\t"+tokenBBoughtAmount.String())

	if !foundA || !foundB {
		return
	}

	go updateTokenUsage(A, B)

	maxProfit := ZERO
	profit := ZERO
	var arbParams ArbParams

	pair1Address := crate2(tokenA, tokenB, skipDex)

	pair1Reserves, ok := finalReserves[pair1Address]
	if !ok {
		return
	}

	var reserveA1 *big.Int
	var reserveB1 *big.Int
	if tokenA.Hash().Big().Cmp(tokenB.Hash().Big()) < 0 {
		reserveA1 = pair1Reserves.ReserveA
		reserveB1 = pair1Reserves.ReserveB
	} else {
		reserveA1 = pair1Reserves.ReserveB
		reserveB1 = pair1Reserves.ReserveA
	}

	// reserves := getDirectPairReserves(tokenA, tokenB, A.Symbol, B.Symbol)
	simulatedReserveA := new(big.Int).Add(reserveA1, tokenASoldAmount)
	simulatedReserveB := new(big.Int).Sub(reserveB1, tokenBBoughtAmount)

	// fmt.Println("reserveB1 ", reserveB1)
	// fmt.Println("sold B ", tokenBBoughtAmount)
	for dexName, dexData := range dexList {
		if dexName != skipDex.Name {
			pair2Address := crate2(tokenA, tokenB, dexData)
			pair2Reserves, ok := finalReserves[pair2Address]

			if !ok {
				continue
			}

			var reserveA2 *big.Int
			var reserveB2 *big.Int
			if tokenA.Hash().Big().Cmp(tokenB.Hash().Big()) < 0 {
				reserveA2 = pair2Reserves.ReserveA
				reserveB2 = pair2Reserves.ReserveB
			} else {
				reserveA2 = pair2Reserves.ReserveB
				reserveB2 = pair2Reserves.ReserveA
			}

			optimalInput := calculateOptimalInput(simulatedReserveB, simulatedReserveA, reserveA2, reserveB2)

			if optimalInput.Cmp(ZERO) > 0 {
				profit = calculateProfit(optimalInput, simulatedReserveB, simulatedReserveA, reserveA2, reserveB2)
				if profit.Cmp(ZERO) <= 0 {
					continue
				}
				// buildLog(&log, "profit "+profit.String()+" "+B.Symbol+" "+dexData.Name)
				if profit.Cmp(new(big.Int).Div(new(big.Int).Mul(optimalInput, big.NewInt(3)), b1000)) > 0 {
					maxProfit = profit
					flashAmount, ok := getAmountOut(optimalInput, simulatedReserveB, simulatedReserveA)
					if !ok {
						return
					}
					arbParams = ArbParams{
						flashToken:   tokenA,
						flashAmount:  flashAmount,
						paybackToken: tokenB,
						path:         []common.Address{pair1Address, pair2Address},
						dex:          dexName,
					}
				}
			}

		}
	}

	minimum := new(big.Int)
	minimum.SetString(B.Minimum, 10)
	// buildLog(&log, "minimum profit: "+B.Minimum+" "+B.Symbol)
	// buildLog(&log, "best profit: "+maxProfit.String()+" "+B.Symbol)
	if maxProfit.Cmp(minimum) > 0 {
		cost := new(big.Int)

		if tx.Type() == 0 {
			cost.Mul(tx.GasPrice(), bn150k).Mul(cost, big.NewInt(4))
		} else if tx.Type() == 2 {
			cost.Mul(tx.GasFeeCap(), bn150k).Mul(cost, big.NewInt(4))
		}

		cost.Mul(cost, TEN.Exp(TEN, big.NewInt(18-B.decimals), nil))
		if cost.Cmp(maxProfit) < 0 {
			go executeArbitrage(arbParams, tx, fromAddress0, privateKey0, nonce0, noticed)
			// go executeArbitrage(arbParams, tx, fromAddress1, privateKey1, nonce1 ,noticed)
			// go executeArbitrage(arbParams, tx, fromAddress2, privateKey2, nonce2 ,noticed)
			// go executeArbitrage(arbParams, tx, fromAddress3, privateKey3, nonce3 ,noticed)
			// go executeArbitrage(arbParams, tx, fromAddress4, privateKey4, nonce4 ,noticed)
			buildLog(&log, "big trade " + tx.Hash().Hex())
			buildLog(&log, "input: "+arbParams.flashAmount.String())
			buildLog(&log, "pair1: "+arbParams.path[0].String())
			buildLog(&log, "pair2: "+arbParams.path[1].String())
			buildLog(&log, "\033[32mprofit: "+maxProfit.String()+" "+B.Symbol+" @ "+arbParams.dex+"\033[0m")
			fmt.Println(log)
		}
	}
}

func executeArbitrage(params ArbParams, tx *types.Transaction, sender common.Address, privateKey *ecdsa.PrivateKey, nonce *big.Int, noticed time.Time) {
	
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal(err)
		return
	}
	auth.Nonce = nonce
	auth.GasLimit = uint64(500_000)
	if tx.Type() == 0 {
		auth.GasPrice = tx.GasPrice()
	} else {
		auth.GasFeeCap = tx.GasFeeCap()
		auth.GasTipCap = tx.GasTipCap()
	}
	auth.Context = context.Background()
	auth.From = sender
	auth.Value = ZERO

	

	start := time.Now()
	transaction, err := arbContract.PrintMoney(auth, params.flashAmount, params.flashToken, params.paybackToken, params.path)
	fmt.Println("tx runTime: ", time.Since(start))
	if err != nil {
		fmt.Println(err)
	}
	nonce.Add(nonce, ONE)

	fmt.Println("transaction ", transaction.Hash().Hex(), "\n", "run time: ", time.Since(noticed))
}
