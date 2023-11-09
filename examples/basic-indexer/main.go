package main

import (
	"context"
	"fmt"
	"log"

	"github.com/alecsavvy/gaudius/gen/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	ctx := context.Background()

	addr := common.HexToAddress("0x1Cd8a543596D499B9b6E7a6eC15ECd2B7857Fd64")
	cl, err := ethclient.Dial("https://acdc-gateway.audius.co/")
	
	if err != nil {
		log.Fatal(err)
	}

	em, err := contracts.NewEntityManager(addr, cl)
	
	if err != nil {
		log.Fatal(err)
	}

	start := uint64(20605697)
	current := start

	for true {
		opts := &bind.FilterOpts{
			Start: current,
			End: &current,
			Context: ctx,
		}
	
		eventiter, err := em.FilterManageEntity(opts)
		defer eventiter.Close()
	
		if err != nil {
			log.Fatal(err)
		}
	
		for eventiter.Next() {
			event := eventiter.Event
			fmt.Println(event)
		}

		current += 1
	}
}