package gaudius

import (
	"fmt"

	"github.com/alecsavvy/gaudius/gen/discovery/models"
)

func (sdk *AudiusSdk) GetUser(id string) (*models.User, error) {
	dn := sdk.Discovery
	res, err := dn.discoveryClient.R().Get(fmt.Sprintf("/users/%s", id))
	if err != nil {
		return nil, err
	}
	var user models.UserResponse
	err = user.UnmarshalBinary(res.Body())
	if err != nil {
		return nil, err
	}

	return user.Data, nil
}

func (sdk *AudiusSdk) getUserFromWallet(wallet string) (*models.User, error) {
	dn := sdk.Discovery
	res, err := dn.discoveryClient.R().Get(fmt.Sprintf("/users/id?associated_wallet=%s", wallet))
	if err != nil {
		return nil, err
	}
	var user models.UserResponse
	err = user.UnmarshalBinary(res.Body())
	if err != nil {
		return nil, err
	}

	return user.Data, nil
}

func (sdk *AudiusSdk) GetUserHandle(handle string) (*models.User, error) {
	dn := sdk.Discovery
	res, err := dn.discoveryClient.R().Get(fmt.Sprintf("/users/handle/%s", handle))
	if err != nil {
		return nil, err
	}
	var user models.UserResponse
	err = user.UnmarshalBinary(res.Body())
	if err != nil {
		return nil, err
	}

	return user.Data, nil
}

func (sdk *AudiusSdk) GetUserAiAttributed(handle string) ([]*models.Track, error) {
	dn := sdk.Discovery
	res, err := dn.discoveryClient.R().Get(fmt.Sprintf("/users/handle/%s/tracks/ai_attributed", handle))
	if err != nil {
		return nil, err
	}
	var user models.TracksResponse
	err = user.UnmarshalBinary(res.Body())
	if err != nil {
		return nil, err
	}

	return user.Data, nil
}

func (sdk *AudiusSdk) UserSearch(query string) ([]*models.User, error) {
	dn := sdk.Discovery
	res, err := dn.discoveryClient.R().Get(fmt.Sprintf("/users/search?query=%s", query))
	if err != nil {
		return nil, err
	}
	var user models.UserSearch
	err = user.UnmarshalBinary(res.Body())
	if err != nil {
		return nil, err
	}

	return user.Data, nil
}

func (sdk *AudiusSdk) GetUserAuthorizedApps(id string) ([]*models.AuthorizedApp, error) {
	dn := sdk.Discovery
	res, err := dn.discoveryClient.R().Get(fmt.Sprintf("/users/%s/authorized_apps", id))
	if err != nil {
		return nil, err
	}
	var apps models.AuthorizedApps
	err = apps.UnmarshalBinary(res.Body())
	if err != nil {
		return nil, err
	}

	return apps.Data, nil
}
