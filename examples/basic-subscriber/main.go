package main

import (
	"fmt"

	"github.com/alecsavvy/gaudius"
	"github.com/alecsavvy/gaudius/gen/contracts"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	sdk := gaudius.NewSdkUnsafe()
	scanner, stopper := sdk.EventSubscriber()

	events := []*contracts.EntityManagerManageEntity{}
	for event := range scanner {
		events = append(events, event)
		if len(events) == 5{
			break
		}
	}
	close(stopper)
	
	fmt.Println("received five events!")
	spew.Dump(events)
}