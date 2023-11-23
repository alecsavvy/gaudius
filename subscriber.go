package gaudius

import (
	"context"
	"math/big"
	"time"

	"github.com/alecsavvy/gaudius/gen/contracts"
	"github.com/ethereum/go-ethereum/common"
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

// method to receive events from a range of blocks (and/or future events)
func (sdk *AudiusSdk) EventScanner(params *ScannerParams) (<-chan *contracts.EntityManagerManageEntity, chan bool) {
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
	stopSignal := make(chan bool)

	go func() {
		defer close(ch)
		current := int64(start)
		select {
		case <-stopSignal:
			return
		default:
			for {
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
		}
	}()
	return ch, stopSignal
}

// method to receive live stream of all current events
func (sdk *AudiusSdk) EventSubscriber() (<-chan *contracts.EntityManagerManageEntity, chan bool) {
	return sdk.EventScanner(nil)
}

// method that listens to current events until a certain transaction hash is found
func (sdk *AudiusSdk) TxSubscriber(sp *ScannerParams, tx common.Hash) <-chan *contracts.EntityManagerManageEntity {
	ch := make(chan *contracts.EntityManagerManageEntity)
	scanner, stopper := sdk.EventScanner(sp)

	go func() {
		// currently we only have one log per tx
		for event := range scanner {
			if event.Raw.TxHash == tx {
				close(stopper)
				ch <- event
				close(ch)
			}
		}
	}()

	return ch
}
