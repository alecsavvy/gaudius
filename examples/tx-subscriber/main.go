package main

import (
	"github.com/alecsavvy/gaudius"
	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/common"
)

func main() {
	sdk := gaudius.NewSdkUnsafe()
	
	tx := common.HexToHash("0xf745d137e31d9b3650b49d2af970951e80e158b4c2e8445b0404e3caccabd978")
	sb := uint64(21838580)
	sp := &gaudius.ScannerParams{ StartBlock: &sb }

	sub := sdk.TxSubscriber(sp, tx)

	event := <-sub
	spew.Dump(event)
}