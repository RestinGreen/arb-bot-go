package main

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

func simulateSwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, reserves []PairReserve) ([]*big.Int, bool) {
	amounts, ok := getAmountsOut(amountIn, path, reserves)
	if amounts[len(amounts)-1].Cmp(amountOutMin) < 0 {
		return amounts, false
	}
	return amounts, ok
}

func simulateSwapTokensForExactTokens(amountOut *big.Int, amountInMax *big.Int, path []common.Address, reserves []PairReserve) ([]*big.Int, bool) {
	amounts, ok := getAmountsIn(amountOut, path, reserves)	
	if amounts[0].Cmp(amountInMax) > 0 {
		return amounts, false
	}
	return amounts, ok
}

func simulateSwapTokensForExactETH(amountOut *big.Int, amountInMax *big.Int, path []common.Address, reserves []PairReserve) ([]*big.Int, bool) {
	amounts, ok := getAmountsIn(amountOut, path, reserves)
	if amounts[0].Cmp(amountInMax) > 0 {
		return amounts, false
	}
	return amounts, ok
}

func simulateSwapExactTokensForETH(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, reserves []PairReserve) ([]*big.Int, bool) {
	amounts, ok := getAmountsOut(amountIn, path, reserves)
	if amounts[len(amounts)-1].Cmp(amountOutMin) < 0 {
		return amounts, false
	}
	return amounts, ok
}

func simulateSwapExactETHForTokens(msgValue *big.Int, amountOutMin *big.Int, path []common.Address, reserves []PairReserve) ([]*big.Int, bool) {
	amounts, ok := getAmountsOut(msgValue, path, reserves)
	if amounts[len(amounts)-1].Cmp(amountOutMin) < 0 {
		return amounts, false
	}
	return amounts, ok
}

func simulateSwapETHForExactTokens(amountOut *big.Int, path []common.Address, reserves []PairReserve) ([]*big.Int, bool) {
	return getAmountsIn(amountOut, path, reserves)

}

func simulateSwapExactETHForTokensSupportingFeeOnTransferTokens() {

}

func simulateSwapExactTokensForETHSupportingFeeOnTransferTokens() {

}

func simulateSwapExactTokensForTokensSupportingFeeOnTransferTokens() {

}