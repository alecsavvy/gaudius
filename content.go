package gaudius

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
)

// type that manages the selected content node
type ContentNode struct {
	ContentNodes []string
	SelectedNode string
	client       *resty.Client
}

func NewContentNode(nodes []string) (*ContentNode, error) {
	selectedNode, err := SelectHealthyContentNode(nodes)
	if err != nil {
		return nil, err
	}
	contentBaseUrl := selectedNode
	contentClient := resty.New().SetBaseURL(contentBaseUrl)
	return &ContentNode{ContentNodes: nodes, SelectedNode: selectedNode, client: contentClient}, nil
}

func SelectHealthyContentNode(contentNodes []string) (string, error) {
	shuffle(contentNodes)
	client := resty.New().SetTimeout(time.Second * 3).GetClient()
	for _, endpoint := range contentNodes {
		route := fmt.Sprintf("%s/health_check", endpoint)
		res, err := client.Get(route)
		if err != nil {
			continue
		}
		if res.StatusCode == 200 {
			return endpoint, nil
		}
	}
	return "", errors.New("found no healthy content nodes")
}
