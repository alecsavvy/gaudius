package gaudius

import (
	"alecsavvy/gaudius/gen/discovery/models"
	"fmt"
)

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
