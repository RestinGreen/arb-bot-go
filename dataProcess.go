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
		
		go checkArbitrage(input.Path[i], input.Path[i+1], output.Amounts[i], output.Amounts[i+1], tx, data.Name)
	}


}

func swapWithSimpleFeeType2(tx *types.Transaction, method *abi.Method, log string, data DexData) {
	buildLog(&log, method.Name)


}

func checkArbitrage(tokenA common.Address, tokenB common.Address, tokenASoldAmount *big.Int, tokenBSoldAmount *big.Int, tx *types.Transaction, skipDex string) {

	A, foundA := getTokenDataByAddress(tokenA.Hex())
	B, foundB := getTokenDataByAddress(tokenB.Hex())
	fmt.Println(A.Symbol)
	fmt.Println(B.Symbol)

	if !foundA || !foundB {
		return
	}

	// maxProfit := ZERO
	// profit := ZERO
	reserves := getDirectPairReserves(tokenA, tokenB, A.Symbol, B.Symbol)
	reserveA, reserveB := decodeStorageSlot(reserves[strings.ToUpper(A.Symbol+B.Symbol+skipDex)]["storage"], tokenA, tokenB)
	fmt.Println(reserveA)
	fmt.Println(reserveB)

}
