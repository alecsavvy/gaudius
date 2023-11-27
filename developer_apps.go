package gaudius

import (
	"fmt"

	"github.com/alecsavvy/gaudius/gen/discovery/models"
)

func (sdk *AudiusSdk) GetDeveloperApp(address string) (*models.DeveloperApp, error) {
	dn := sdk.Discovery
	res, err := dn.discoveryClient.R().Get(fmt.Sprintf("/developer_apps/%s", address))
	if err != nil {
		return nil, err
	}
	var resp models.DeveloperAppResponse
	err = resp.UnmarshalBinary(res.Body())
	if err != nil {
		return nil, err
	}

	return resp.Data, nil
}

func (sdk *AudiusSdk) GetDeveloperApps(userID string) ([]*models.DeveloperApp, error) {
	dn := sdk.Discovery
	res, err := dn.discoveryClient.R().Get(fmt.Sprintf("/users/%s/developer_apps", userID))
	if err != nil {
		return nil, err
	}
	var resp models.DeveloperApps
	err = resp.UnmarshalBinary(res.Body())
	if err != nil {
		return nil, err
	}

	return resp.Data, nil
}
