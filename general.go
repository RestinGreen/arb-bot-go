package main

import (
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	wsEndpoint  = ""
	ipcEndpoint = ""
	dexJson     = ""
	uniV2Json   = ""
	ethClient *ethclient.Client	= nil
)

func buildLog(log *string, text string) {
	*log += time.Now().String() + text + "\n"
}