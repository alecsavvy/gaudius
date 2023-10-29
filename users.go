package gaudius

import (
	"alecsavvy/gaudius/gen/discovery/models"
	"fmt"
)

func (sdk *AudiusSdk) getUserHandle(handle string) (*models.User, error) {
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

func (sdk *AudiusSdk) getUserAiAttributed(handle string) ([]*models.Track, error) {
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

func (sdk *AudiusSdk) getUserAssociatedWallet(address string) (*models.EncodedUserID, error) {
	dn := sdk.Discovery
	res, err := dn.discoveryClient.R().Get(fmt.Sprintf("/users/id?associated_wallet=%s", address))
	if err != nil {
		return nil, err
	}
	var user models.UserAssociatedWalletResponse
	err = user.UnmarshalBinary(res.Body())
	if err != nil {
		return nil, err
	}

	return user.Data, nil
}

func (sdk *AudiusSdk) getUser(id string) (*models.User, error) {
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

func (sdk *AudiusSdk) userSearch(query string) ([]*models.User, error) {
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
