package gaudius

import "github.com/go-resty/resty/v2"

// structure that manages the selected discovery node
type DiscoveryNode struct {
	DiscoveryNodes []string
	SelectedNode string
	discoveryClient *resty.Client
	discoveryFullClient *resty.Client
	storageClient *resty.Client
}

func NewDiscoveryNode(selectedNode string) *DiscoveryNode {
	var discoveryNodes []string

	return &DiscoveryNode{ DiscoveryNodes: discoveryNodes, SelectedNode: selectedNode }
}
