package main

import (
	"crypto/ecdsa"
	"main/DoSimpleArb"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/joho/godotenv"
	
)

func initBot() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	wsEndpoint = os.Getenv("NODE_WS_ENDPOINT")
	ipcEndpoint = os.Getenv("NODE_IPC_ENDPOINT")
	dexJson = os.Getenv("DEX_JSON")
	uniV2Json = os.Getenv("UNIV2_JSON")

	readDexFromFile(dexJson)
	readUniV2Router02(uniV2Json)

	user := os.Getenv("DB_USER")
	pass := os.Getenv("PASS")
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	name := os.Getenv("NAME")

	initClient()
	initDB(user, pass, host, port, name)
	initContract()
	initWallets()
	
}

func initClient() {
	var err error

	baseClient, err = rpc.Dial(wsEndpoint)
	if err != nil {
		panic(err)
	}
	ethClient, err = ethclient.Dial(ipcEndpoint)
	if err != nil {
		panic(err)
	}
}

func initWallets() {

	var err error

	address0 := os.Getenv("PRIVATE_KEY0")
	address1 := os.Getenv("PRIVATE_KEY1")
	address2 := os.Getenv("PRIVATE_KEY2")
	address3 := os.Getenv("PRIVATE_KEY3")
	address4 := os.Getenv("PRIVATE_KEY4")
	
	privateKey0, err = crypto.HexToECDSA(address0)
	if err != nil {
		panic(err)
	}
	privateKey1, err = crypto.HexToECDSA(address1)
	if err != nil {
		panic(err)
	}
	privateKey2, err = crypto.HexToECDSA(address2)
	if err != nil {
		panic(err)
	}
	privateKey3, err = crypto.HexToECDSA(address3)
	if err != nil {
		panic(err)
	}
	privateKey4, err = crypto.HexToECDSA(address4)
	if err != nil {
		panic(err)
	}
	
	publicKey0 := privateKey0.Public()	
	publicKey1 := privateKey1.Public()
	publicKey2 := privateKey2.Public()
	publicKey3 := privateKey3.Public()
	publicKey4 := privateKey4.Public()	

	publicKeyECDSA0, ok := publicKey0.(*ecdsa.PublicKey)
	if !ok {
		panic("error casting public key 0 to ECDSA")
	}
	publicKeyECDSA1, ok := publicKey1.(*ecdsa.PublicKey)
	if !ok {
		panic("error casting public key 1 to ECDSA")
	}
	publicKeyECDSA2, ok := publicKey2.(*ecdsa.PublicKey)
	if !ok {
		panic("error casting public key 2 to ECDSA")
	}
	publicKeyECDSA3, ok := publicKey3.(*ecdsa.PublicKey)
	if !ok {
		panic("error casting public key 3 to ECDSA")
	}
	publicKeyECDSA4, ok := publicKey4.(*ecdsa.PublicKey)
	if !ok {
		panic("error casting public key 4 to ECDSA")
	}

	fromAddress0 = crypto.PubkeyToAddress(*publicKeyECDSA0)
	fromAddress1 = crypto.PubkeyToAddress(*publicKeyECDSA1)
	fromAddress2 = crypto.PubkeyToAddress(*publicKeyECDSA2)
	fromAddress3 = crypto.PubkeyToAddress(*publicKeyECDSA3)
	fromAddress4 = crypto.PubkeyToAddress(*publicKeyECDSA4)



}

func initContract() {
	var err error

	contractHexAddress := os.Getenv("CONTRACT_ADDRESS")
	
	contractAddress := common.HexToAddress(contractHexAddress)	
	contract, err = DoSimpleArb.NewDoSimpleArb(contractAddress, ethClient)
	if err != nil {
		panic(err)
	}
}