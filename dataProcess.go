package main

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/mitchellh/mapstructure"
)

type UniV2InputsType1 struct {
	Amount1 *big.Int
	Path []common.Address
	To common.Address
	Deadline *big.Int
}

type UniV2Outputs struct {
	Amounts []*big.Int
}

type UniV2InputsType2 struct {
	Amount1 *big.Int
	Amount2 *big.Int
	Path []common.Address
	To common.Address
	Deadline *big.Int
}

func swapWithSimpleFeeType1(tx *types.Transaction, method *abi.Method, log string, data DexData) {
	buildLog(&log, method.Name)
	fmt.Println(tx.Hash())

	var input UniV2InputsType1
	inputMap := make(map[string]interface{}, 0)
	err := method.Inputs.UnpackIntoMap(inputMap, tx.Data()[4:])
	if err != nil {
		panic(err)
	}
	err = mapstructure.Decode(inputMap, &input)
	if err != nil {
		panic(err)
	}

	sender, err := types.Sender(types.NewEIP155Signer(tx.ChainId()), tx)
	if err != nil {
		sender, err = types.Sender(types.NewLondonSigner(tx.ChainId()), tx)
	}
	if err != nil {
		fmt.Println(err)
		return
	}

	callMsg := ethereum.CallMsg{
		To: tx.To(),
		Data: tx.Data(),
		From: sender,
	}

	start := time.Now()
	simulation, err := ethClient.CallContract(context.Background(), callMsg, nil)
	fmt.Println(time.Since(start))
	if err != nil {
		fmt.Println(err)
		return
	}
	var output UniV2Outputs
	outputMap := make(map[string]interface{}, 0)
	err = method.Outputs.UnpackIntoMap(outputMap, simulation)
	if err != nil {
		panic (err)
	}

	err = mapstructure.Decode(outputMap, &output)
	if err != nil {
		panic(err)
	}


	for i:= 0; i < len(input.Path)-1; i++ {
		
		go checkArbitrage(input.Path[i], input.Path[i+1], output.Amounts[i], output.Amounts[i+1], tx, data.Name, log)
	}


}

func swapWithSimpleFeeType2(tx *types.Transaction, method *abi.Method, log string, data DexData) {
	buildLog(&log, method.Name)


}

func checkArbitrage(tokenA common.Address, tokenB common.Address, tokenASoldAmount *big.Int, tokenBBoughtAmount *big.Int, tx *types.Transaction, skipDex string, log string) {

	A, foundA := getTokenDataByAddress(tokenA.Hex())
	B, foundB := getTokenDataByAddress(tokenB.Hex())
	fmt.Println(A.Symbol)
	fmt.Println(B.Symbol)

	if !foundA || !foundB {
		return
	}

	maxProfit := ZERO
	profit := ZERO
	var arbParams ArbParams 

	reserves := getDirectPairReserves(tokenA, tokenB, A.Symbol, B.Symbol)
	reserveA1, reserveB1 := decodeStorageSlot(reserves[strings.ToUpper(A.Symbol+B.Symbol+skipDex)]["storage"], tokenA, tokenB)
	pair1 := reserves[strings.ToUpper(A.Symbol+B.Symbol+skipDex)]["address"]
	simulatedReserveA := new(big.Int).Add(reserveA1, tokenASoldAmount)
	simulatedReserveB := new(big.Int).Sub(reserveB1, tokenBBoughtAmount)

	for dex, _ := range dexList {
		if dex != skipDex {
			reserveA2, reserveB2 := decodeStorageSlot(reserves[strings.ToUpper(A.Symbol+B.Symbol+dex)]["storage"], tokenA, tokenB)
			pair2 := reserves[strings.ToUpper(A.Symbol+B.Symbol+dex)]["address"]

			optimalInput := calculateOptimalInput(simulatedReserveB, simulatedReserveA, reserveA2, reserveB2)

			if optimalInput.Cmp(ZERO) > 0 {
				profit = calculateProfit(optimalInput, simulatedReserveB, simulatedReserveA, reserveA2, reserveB2)
				buildLog(&log, "profit " + profit.String() + " " + B.Symbol + " " + dex)
				if profit.Cmp( new(big.Int).Div(new(big.Int).Mul(optimalInput, big.NewInt(3)), b1000)    ) > 0 {
					maxProfit = profit
					flashAmount := getAmountOut(*simulatedReserveB, *simulatedReserveA, *optimalInput)
					arbParams = ArbParams{
						flashToken: tokenA,
						flashAmount: flashAmount,
						paybackToken: tokenB,
						paybackAmount: optimalInput,
						path: []common.Address {common.HexToAddress(pair1), common.HexToAddress(pair2)},
					}
				}
			}

		}
	}

	minimum := new(big.Int)
	minimum.SetString(B.Minimum, 10)
	buildLog(&log, "minimum profit: " + B.Minimum + " " + B.Symbol)
	if maxProfit.Cmp(minimum) > 0 {
		if tx.Type() == 0 {
			cost := new(big.Int)
			cost.Mul(tx.GasPrice(), bn150k).Mul(cost, big.NewInt(4))
			cost.Mul(cost, TEN.Exp(TEN, big.NewInt(18-B.decimals), nil))
			if cost.Cmp(maxProfit) < 0 {
				go executeArbitrage(arbParams, tx, )
			}
		} else if tx.Type() == 2 {

		}
	}
}


func executeArbitrage(params ArbParams, tx *types.Transaction, sender common.Address)