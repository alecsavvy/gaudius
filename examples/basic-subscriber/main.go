package main

import (
	"fmt"

	"github.com/alecsavvy/gaudius"
)

func main() {
	sdk := gaudius.NewSdkUnsafe()
	scanner := sdk.EventSubscriber()
	for event := range scanner {
		fmt.Println(event)
	}
}