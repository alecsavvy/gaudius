package main

import (
	"fmt"
	"log"
	"os"

	"github.com/alecsavvy/gaudius"
	"github.com/alecsavvy/gaudius/gen/contracts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	url := os.Getenv("ETHEREUM_RPC")
	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatal(err)
	}

	addr := common.HexToAddress(gaudius.MainnetRegistryAddress)
	contractRegistry, err := contracts.NewRegistry(addr, client)
	if err != nil {
		log.Fatal(err)
	}

	audTokenAddr, err := contractRegistry.GetContract0(nil, gaudius.AudiusTokenKey)
	if err != nil {
		log.Fatal(err)
	}

	audioToken, err := contracts.NewAudiusToken(audTokenAddr, client)
	if err != nil {
		log.Fatal(err)
	}

	totalSupply, err := audioToken.TotalSupply(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Total audio token supply: %d\n", totalSupply)
}
