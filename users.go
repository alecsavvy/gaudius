package gaudius

import (
	"alecsavvy/gaudius/gen/discovery/models"
	"fmt"
)

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