package main

import (
	"fmt"
	"log"
	"os"

	"github.com/alecsavvy/gaudius"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	ethrpc := os.Getenv("ETHEREUM_RPC_URL")

	params := gaudius.NewAudiusSdkMainnetParams()
	params.EthereumRpc = ethrpc

	sdk, err := gaudius.NewCustomSdk(params)
	if err != nil {
		log.Fatal(err)
	}

	tokenContract, err := sdk.Contracts.GetAudioTokenContract()
	if err != nil {
		log.Fatal(err)
	}

	totalSupply, err := tokenContract.TotalSupply(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Total $AUDIO token supply: %d", totalSupply)
}
