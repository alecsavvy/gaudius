package gaudius

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
)

type StorageNodeConfig struct {
	Endpoint string

	storageClient         *resty.Client
	storageInternalClient *resty.Client
}

// type that manages the selected storage node
type StorageNode struct {
	StorageNodes []string
	SelectedNode string
	client       *resty.Client
}

func NewStorageNode(nodes []string) (*StorageNode, error) {
	selectedNode, err := SelectHealthyStorageNode(nodes)
	if err != nil {
		return nil, err
	}
	storageBaseUrl := selectedNode
	storageClient := resty.New().SetBaseURL(storageBaseUrl)
	return &StorageNode{StorageNodes: nodes, SelectedNode: selectedNode, client: storageClient}, nil
}

func SelectHealthyStorageNode(storageNodes []string) (string, error) {
	shuffle(storageNodes)
	client := resty.New().SetTimeout(time.Second * 3).GetClient()
	for _, endpoint := range storageNodes {
		route := fmt.Sprintf("%s/health_check", endpoint)
		res, err := client.Get(route)
		if err != nil {
			continue
		}
		if res.StatusCode == 200 {
			return endpoint, nil
		}
	}
	return "", errors.New("found no healthy discovery nodes")
}
