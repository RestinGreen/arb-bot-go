package main

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	wsEndpoint  = ""
	ipcEndpoint = ""
	dexJson     = ""
	uniV2Json   = ""
	ethClient *ethclient.Client	= nil

	ZERO *big.Int = big.NewInt(0)
	ONE *big.Int = big.NewInt(1)
	b997 *big.Int = big.NewInt(997)
	b1000 *big.Int = big.NewInt(1000)
)

func buildLog(log *string, text string) {
	*log += time.Now().String() + text + "\n"
}

func getAmountOut(reserveIn big.Int, reserveOut big.Int, amountIn big.Int) *big.Int {
	amountInWithFee := new(big.Int).Mul(&amountIn, b997)
	numerator := new(big.Int).Mul(amountInWithFee, &reserveOut)
	denominator := new(big.Int)
	denominator.Mul(&reserveIn, b1000).Add(denominator, amountInWithFee)

	return numerator.Div(numerator, denominator)
}

func getAmountIn(reserveIn big.Int, reserveOut big.Int, amountOut big.Int) *big.Int {
	numerator := new(big.Int)
	numerator.Mul(&reserveIn, &amountOut).Mul(numerator, b1000)
	denominator := new(big.Int)
	denominator.Sub(&reserveOut, &amountOut).Mul(denominator, b997)

	result := new(big.Int)
	result.Div(numerator, denominator).Add(result, ONE)
	return result
}

func sortAddtess(tokenA common.Address, tokenB common.Address) (common.Address, common.Address) {

	if tokenA.Hash().Big().Cmp(tokenB.Hash().Big()) < 0 {
		return tokenA, tokenB
	} else {
		return tokenB, tokenA
	}
}

func crate2(tokenA common.Address, tokenB common.Address, data DexData) common.Address {

	tA, tB := sortAddtess(tokenA, tokenB)
	ff := []byte{255}
	ff = append(ff, data.Factory.Bytes()...)
	ff = append(ff, crypto.Keccak256(append(tA.Bytes(), tB.Bytes()...))...)
	ff = append(ff, common.Hex2Bytes(data.Salt)...)

	return common.BytesToAddress(crypto.Keccak256(ff))
}

func decodeStorageSlot(storage string, tokenA common.Address, tokenB common.Address) (*big.Int, *big.Int) {
	
	// fmt.Println("storage ", storage)
	base := storage[2:]
	// fmt.Println("base ", base)

	reserve0 := base[36 : 64]
	// fmt.Println("reserve0 ", reserve0)
	r0, _ := new(big.Int).SetString(reserve0, 16)
	// fmt.Println("r0 ", r0)
	reserve1 := base[8 : 36]
	r1, _:= new(big.Int).SetString(reserve1, 16)

	if tokenA.Hash().Big().Cmp(tokenA.Hash().Big()) < 0 {
		return  r0, r1
	} else {
		return r1, r0
	}


}