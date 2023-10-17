package gaudius

// structure that manages the selected discovery node
type DiscoveryNode struct {
	DiscoveryNodes []string
	SelectedNode string
}

func NewDiscoveryNode(selectedNode *string) *DiscoveryNode {
	if (selectedNode != nil) {
		var discoveryNodes []string
		return &DiscoveryNode{ DiscoveryNodes: discoveryNodes, SelectedNode: *selectedNode }
	}
	var discoveryNodes []string
	return &DiscoveryNode{ DiscoveryNodes: discoveryNodes, SelectedNode: *selectedNode }
}
