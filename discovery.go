package gaudius

import "github.com/go-resty/resty/v2"

// structure that manages the selected discovery node
type DiscoveryNode struct {
	DiscoveryNodes []string
	SelectedNode string
	client *resty.Client
}

func NewDiscoveryNode(selectedNode *string) *DiscoveryNode {
	if (selectedNode != nil) {
		var discoveryNodes []string
		return &DiscoveryNode{ DiscoveryNodes: discoveryNodes, SelectedNode: *selectedNode }
	}
	var discoveryNodes []string
	return &DiscoveryNode{ DiscoveryNodes: discoveryNodes, SelectedNode: "https://discoveryprovider3.audius.co" }
}
