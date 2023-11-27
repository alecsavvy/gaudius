package gaudius

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

// structure that manages the selected discovery node
type DiscoveryNode struct {
	DiscoveryNodes []string
	SelectedNode   string

	discoveryClient     *resty.Client
	discoveryFullClient *resty.Client
}

func NewDiscoveryNode(selectedNode string) *DiscoveryNode {
	var discoveryNodes []string
	discoveryBaseUrl := fmt.Sprintf("%s/v1", selectedNode)
	discoveryFullBaseUrl := fmt.Sprintf("%s/full", discoveryBaseUrl)
	discoveryClient := resty.New().SetBaseURL(discoveryBaseUrl)
	discoveryFullClient := resty.New().SetBaseURL(discoveryFullBaseUrl)
	return &DiscoveryNode{DiscoveryNodes: discoveryNodes, SelectedNode: selectedNode, discoveryClient: discoveryClient, discoveryFullClient: discoveryFullClient}
}
