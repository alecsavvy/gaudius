package main

import (
	"fmt"
	"log"

	"github.com/alecsavvy/gaudius"
)

func main() {
	sdk, err := gaudius.NewSdk()
	if err != nil {
		log.Fatal(err)
	}

	selectedDiscovery := sdk.Discovery.SelectedNode
	selectedContent := sdk.Storage.SelectedNode

	fmt.Printf("Discovery: %s\nContent: %s\n", selectedDiscovery, selectedContent)
}
