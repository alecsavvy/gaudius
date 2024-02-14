package gaudius

import (
	"errors"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-resty/resty/v2"
)

// structure that manages the selected discovery node
type DiscoveryNode struct {
	DiscoveryNodes []string
	SelectedNode   string

	discoveryClient     *resty.Client
	discoveryFullClient *resty.Client

	AcdcClient *ethclient.Client
}

func NewDiscoveryNode(nodes []string) (*DiscoveryNode, error) {
	selectedNode, err := SelectHealthyDiscoveryNode(nodes)
	if err != nil {
		return nil, err
	}
	discoveryBaseUrl := fmt.Sprintf("%s/v1", selectedNode)
	discoveryFullBaseUrl := fmt.Sprintf("%s/full", discoveryBaseUrl)
	discoveryClient := resty.New().SetBaseURL(discoveryBaseUrl)
	discoveryFullClient := resty.New().SetBaseURL(discoveryFullBaseUrl)
	return &DiscoveryNode{DiscoveryNodes: nodes, SelectedNode: selectedNode, discoveryClient: discoveryClient, discoveryFullClient: discoveryFullClient}, nil
}

func SelectHealthyDiscoveryNode(discoveryNodes []string) (string, error) {
	shuffle(discoveryNodes)
	client := resty.New().SetTimeout(time.Second * 3).GetClient()
	for _, endpoint := range discoveryNodes {
		route := fmt.Sprintf("%s/health_check?enforce_block_diff=true&healthy_block_diff=250&plays_count_max_drift=720", endpoint)
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
