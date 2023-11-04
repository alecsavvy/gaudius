package gaudius

import (
	"encoding/json"
	"fmt"

	"github.com/AudiusProject/audius-protocol/mediorum/server"
)

func (sdk *AudiusSdk) GetUploads() (*[]server.Upload, error) {
	cn := sdk.Storage

	res, err := cn.client.R().Get("/uploads")
	if err != nil {
		return nil, err
	}

	uploads := new([]server.Upload)
	err = json.Unmarshal(res.Body(), uploads)
	if err != nil {
		return nil, err
	}

	return uploads, nil
}

func (sdk *AudiusSdk) GetUpload(id string) (*server.Upload, error) {
	cn := sdk.Storage

	path := fmt.Sprintf("/uploads/%s", id)
	res, err := cn.client.R().Get(path)
	if err != nil {
		return nil, err
	}

	uploads := new(server.Upload)
	err = json.Unmarshal(res.Body(), uploads)
	if err != nil {
		return nil, err
	}

	return uploads, nil
}
