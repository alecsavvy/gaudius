package gaudius

import (
	"context"
	"math/big"
	"time"

	"github.com/alecsavvy/gaudius/gen/contracts"
)

// params for the subscriber function
type ScannerParams struct {
	// block to start with
	// leave empty to start at current block
	StartBlock *uint64
	// block to end subscription
	// leave empty to continue endlessly
	EndBlock *uint64
}

// helper for getting a live subscription of events
func (sdk *AudiusSdk) EventSubscriber() <-chan *contracts.EntityManagerManageEntity {
	return sdk.EventScanner(nil)
}

// configurable event scanner
func (sdk *AudiusSdk) EventScanner(params *ScannerParams) <-chan *contracts.EntityManagerManageEntity {
	ctx := context.TODO()

	// set start block
	var start uint64
	if params != nil && params.StartBlock != nil {
		start = uint64(*params.StartBlock)
	} else {
		current, _ := sdk.AcdcClient.BlockByNumber(ctx, nil)
		start = current.NumberU64() - 1
	}

	ch := make(chan *contracts.EntityManagerManageEntity)

	go func() {
		current := int64(start)
		for true {
			block, err := sdk.AcdcClient.BlockByNumber(ctx, big.NewInt(current))
			if err != nil && err.Error() == "not found" || block == nil {
				// give a few seconds for main chain
				// to get ahead of scanner
				<- await(2 * time.Second)
				continue
			}

			txs := block.Transactions()
		
			for _, tx := range txs {
				if tx == nil {
					continue
				}
				receipt, _ := sdk.AcdcClient.TransactionReceipt(ctx, tx.Hash())
				logs := receipt.Logs
				for _, log := range logs {
					event, _ := sdk.EntityManager.ParseManageEntity(*log)
					ch <- event
				}
			}
			// offset next block being mined
			<- await(500 * time.Millisecond)
			current += 1
		}
	}()
	return ch
}

func await(d time.Duration) <-chan bool {
	c := make(chan bool)

	go func() {
		time.Sleep(d)
		c <- true
		close(c)
	}()

	return c
}
