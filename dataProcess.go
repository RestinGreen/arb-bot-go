package main

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/mitchellh/mapstructure"
)

type Univ2InputsType1 struct {
	Amount1 *big.Int
	Path []common.Address
	To common.Address
	Deadline *big.Int
}

type Univ2InputsType2 struct {
	Amount1 *big.Int
	Amount2 *big.Int
	Path []common.Address
	To common.Address
	Deadline *big.Int
}

func swapWithSimpleFeeType1(tx *types.Transaction, method *abi.Method, log string) {
	buildLog(&log, method.Name)

	var input Univ2InputsType1
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

	simulation, err := ethClient.CallContract(context.Background(), callMsg, nil)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(simulation)
	// for i:= 0; i < len(input.Path); i++ {
	// 	fmt.Println(input.Path[i])
	// }


}

func swapWithSimpleFeeType2(tx *types.Transaction, method *abi.Method, log string) {
	buildLog(&log, method.Name)


}