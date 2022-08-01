package main

import (
	"os"
)

var uniV2ABIString string

func readUniV2Router02(fileName string) {
	fileBytes, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	uniV2ABIString = string(fileBytes)
}
