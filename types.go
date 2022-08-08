package main

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type ArbParams struct {
	flashToken 		common.Address
	paybackToken 	common.Address
	flashAmount		*big.Int
	paybackAmount	*big.Int
	path			[]common.Address
}