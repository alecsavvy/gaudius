package gaudius

import (
	"github.com/AudiusProject/audius-protocol/mediorum/server"
	"github.com/go-resty/resty/v2"
)

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

func NewStorageNode(selectedNode string) *StorageNode {
	var storageNodes []StorageNodeConfig
	storageBaseUrl := selectedNode
	storageClient := resty.New().SetBaseURL(storageBaseUrl)
	return &StorageNode{ StorageNodes: storageNodes, SelectedNode: &StorageNodeConfig{ Endpoint: selectedNode, DelegateOwnerWallet: ""}, client: storageClient}
}

// keeping for ref in audius-protocol pr
func (sn *StorageNode) getUpload() *server.Upload {
	upload := &server.Upload{}
	return upload
}
