package gaudius

import "github.com/go-resty/resty/v2"

type StorageNodeConfig struct {
	Endpoint string
	DelegateOwnerWallet string

	storageClient *resty.Client
	storageInternalClient *resty.Client
}

// type that manages the selected storage node
type StorageNode struct {
	StorageNodes []StorageNodeConfig
	SelectedNode *StorageNodeConfig
	client *resty.Client
}

func NewStorageNode() *StorageNode {
	var storageNodes []StorageNodeConfig
	return &StorageNode{ StorageNodes: storageNodes, SelectedNode: &StorageNodeConfig{ Endpoint: "https://creatornode.audius.co", DelegateOwnerWallet: "0xc8d0C29B6d540295e8fc8ac72456F2f4D41088c8"} }
}
