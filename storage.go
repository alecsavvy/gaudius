package gaudius

// type that manages the selected storage node
type StorageNode struct {
	StorageNodes []string
	SelectedNode string
}

func NewStorageNode(selectedNode *string) *StorageNode {
	if (selectedNode != nil) {
		var storageNodes []string
		return &StorageNode{ StorageNodes: storageNodes, SelectedNode: *selectedNode }
	}
	var storageNodes []string
	return &StorageNode{ StorageNodes: storageNodes, SelectedNode: *selectedNode }
}

