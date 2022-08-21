package main

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

func simulateSwapExactTokensForTokens(amountIn *big.Int, path []common.Address, reserves []PairReserve) ([]*big.Int, bool) {
	return getAmountsOut(amountIn, path, reserves)
}

func simulateSwapTokensForExactTokens(amountOut *big.Int, path []common.Address, reserves []PairReserve) ([]*big.Int, bool) {
	return getAmountsIn(amountOut, path, reserves)	
}

func simulateSwapTokensForExactETH(amountOut *big.Int, path []common.Address, reserves []PairReserve) ([]*big.Int, bool) {
	return getAmountsIn(amountOut, path, reserves)
}

func simulateSwapExactTokensForETH(amountIn *big.Int, path []common.Address, reserves []PairReserve) ([]*big.Int, bool) {
	return getAmountsOut(amountIn, path, reserves)
}

func simulateSwapExactETHForTokens(msgValue *big.Int, path []common.Address, reserves []PairReserve) ([]*big.Int, bool) {
	return getAmountsOut(msgValue, path, reserves)
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