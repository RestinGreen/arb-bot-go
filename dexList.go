package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/ethereum/go-ethereum/common"
)

type DexType int

const (
	UniV2 DexType = iota
	UniV3
	Curve
)

type DexData struct {
	Router  common.Address
	Factory common.Address
	Salt    string 
	DexType DexType
}

var dexList = map[string]DexData{}

func readDexFromFile(jsonFileName string) {
	
	type TempDexData struct {
		Router  string `json:"router"`
		Factory string `json:"factory"`
		Salt    string `json:"salt"`
		DexType string `json:"dexType"`
	}

	jsonFile, err := os.Open(jsonFileName)
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		panic(err)
	}
	var temp map[string]TempDexData
	json.Unmarshal([]byte(byteValue), &temp)


	for key, value := range temp {
		var dexType DexType
		if value.DexType == "UniV2" {
			dexType = UniV2
		} else if value.DexType == "UniV3" {
			dexType = UniV3
		} else {
			dexType = Curve
		}
		dexList[key] = DexData {
			Router: common.HexToAddress(value.Router),
			Factory: common.HexToAddress(value.Factory),
			Salt: value.Salt,
			DexType: dexType,
		}
	}
}
