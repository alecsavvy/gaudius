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
	var resp models.UserResponse
	err = resp.UnmarshalBinary(res.Body())
	if err != nil {
		return nil, err
	}

	return resp.Data, nil
}

func (sdk *AudiusSdk) GetUserFromWallet(wallet string) (*models.User, error) {
	dn := sdk.Discovery
	res, err := dn.discoveryClient.R().Get(fmt.Sprintf("/users/id?associated_wallet=%s", wallet))
	if err != nil {
		return nil, err
	}
	var resp models.UserResponse
	err = resp.UnmarshalBinary(res.Body())
	if err != nil {
		return nil, err
	}

	return resp.Data, nil
}

func (sdk *AudiusSdk) UserSearch(query string) ([]*models.User, error) {
	dn := sdk.Discovery
	res, err := dn.discoveryClient.R().Get(fmt.Sprintf("/users/search?query=%s", query))
	if err != nil {
		return nil, err
	}
	var resp models.UserSearch
	err = resp.UnmarshalBinary(res.Body())
	if err != nil {
		return nil, err
	}

	return resp.Data, nil
}

func (sdk *AudiusSdk) VerifyToken(token string) (*models.DecodedUserToken, error) {
	dn := sdk.Discovery
	res, err := dn.discoveryClient.R().Get(fmt.Sprintf("/users/verify_token?token=%s", token))
	if err != nil {
		return nil, err
	}
	var resp models.VerifyToken
	err = resp.UnmarshalBinary(res.Body())
	if err != nil {
		return nil, err
	}

	return resp.Data, nil
}

func (sdk *AudiusSdk) GetUserConnectedWallets(id string) (*models.ConnectedWallets, error) {
	dn := sdk.Discovery
	res, err := dn.discoveryClient.R().Get(fmt.Sprintf("/users/%s/connected_wallets", id))
	if err != nil {
		return nil, err
	}
	var resp models.ConnectedWalletsResponse
	err = resp.UnmarshalBinary(res.Body())
	if err != nil {
		return nil, err
	}

	return resp.Data, nil
}

func (sdk *AudiusSdk) GetUserFavorites(id string) ([]*models.Favorite, error) {
	dn := sdk.Discovery
	res, err := dn.discoveryClient.R().Get(fmt.Sprintf("/users/%s/favorites", id))
	if err != nil {
		return nil, err
	}
	var resp models.FavoritesResponse
	err = resp.UnmarshalBinary(res.Body())
	if err != nil {
		return nil, err
	}

	return resp.Data, nil
}

func (sdk *AudiusSdk) GetUserReposts(id string, query map[string]string) ([]*models.Activity, error) {
	dn := sdk.Discovery
	res, err := dn.discoveryClient.R().SetQueryParams(query).Get(fmt.Sprintf("/users/%s/reposts", id))
	if err != nil {
		return nil, err
	}
	var resp models.Reposts
	err = resp.UnmarshalBinary(res.Body())
	if err != nil {
		return nil, err
	}

	return resp.Data, nil
}

func (sdk *AudiusSdk) GetUserFollowers(id string, query map[string]string) ([]*models.User, error) {
	dn := sdk.Discovery
	res, err := dn.discoveryClient.R().SetQueryParams(query).Get(fmt.Sprintf("/users/%s/followers", id))
	if err != nil {
		return nil, err
	}
	var resp models.FollowersResponse
	err = resp.UnmarshalBinary(res.Body())
	if err != nil {
		return nil, err
	}

	return resp.Data, nil
}

func (sdk *AudiusSdk) GetUserFollowing(id string, query map[string]string) ([]*models.User, error) {
	dn := sdk.Discovery
	res, err := dn.discoveryClient.R().SetQueryParams(query).Get(fmt.Sprintf("/users/%s/following", id))
	if err != nil {
		return nil, err
	}
	var resp models.FollowingResponse
	err = resp.UnmarshalBinary(res.Body())
	if err != nil {
		return nil, err
	}

	return resp.Data, nil
}

func (sdk *AudiusSdk) GetUserSupporters(id string, query map[string]string) ([]*models.Supporter, error) {
	dn := sdk.Discovery
	res, err := dn.discoveryClient.R().SetQueryParams(query).Get(fmt.Sprintf("/users/%s/supporters", id))
	if err != nil {
		return nil, err
	}
	var resp models.GetSupporters
	err = resp.UnmarshalBinary(res.Body())
	if err != nil {
		return nil, err
	}

	return resp.Data, nil
}

func (sdk *AudiusSdk) GetUserSupporting(id string, query map[string]string) ([]*models.Supporting, error) {
	dn := sdk.Discovery
	res, err := dn.discoveryClient.R().SetQueryParams(query).Get(fmt.Sprintf("/users/%s/supporting", id))
	if err != nil {
		return nil, err
	}
	var resp models.GetSupporting
	err = resp.UnmarshalBinary(res.Body())
	if err != nil {
		return nil, err
	}

	return resp.Data, nil
}

func (sdk *AudiusSdk) GetUserTracks(id string, query map[string]string) ([]*models.Track, error) {
	dn := sdk.Discovery
	res, err := dn.discoveryClient.R().SetQueryParams(query).Get(fmt.Sprintf("/users/%s/tracks", id))
	if err != nil {
		return nil, err
	}
	var resp models.TracksResponse
	err = resp.UnmarshalBinary(res.Body())
	if err != nil {
		return nil, err
	}

	return resp.Data, nil
}

func (sdk *AudiusSdk) GetRelatedUsers(id string, query map[string]string) ([]*models.User, error) {
	dn := sdk.Discovery
	res, err := dn.discoveryClient.R().SetQueryParams(query).Get(fmt.Sprintf("/users/%s/related", id))
	if err != nil {
		return nil, err
	}
	var resp models.RelatedArtistResponse
	err = resp.UnmarshalBinary(res.Body())
	if err != nil {
		return nil, err
	}

	return resp.Data, nil
}

func (sdk *AudiusSdk) GetUserSubscribers(id string, query map[string]string) ([]*models.User, error) {
	dn := sdk.Discovery
	res, err := dn.discoveryClient.R().SetQueryParams(query).Get(fmt.Sprintf("/users/%s/subscribers", id))
	if err != nil {
		return nil, err
	}
	var resp models.SubscribersResponse
	err = resp.UnmarshalBinary(res.Body())
	if err != nil {
		return nil, err
	}

	return resp.Data, nil
}

// most used tags
func (sdk *AudiusSdk) GetUserTags(id string, query map[string]string) ([]string, error) {
	dn := sdk.Discovery
	res, err := dn.discoveryClient.R().SetQueryParams(query).Get(fmt.Sprintf("/users/%s/tags", id))
	if err != nil {
		return nil, err
	}
	var resp models.TagsResponse
	err = resp.UnmarshalBinary(res.Body())
	if err != nil {
		return nil, err
	}

	return resp.Data, nil
}

func (sdk *AudiusSdk) GetUserByHandle(handle string) (*models.User, error) {
	dn := sdk.Discovery
	res, err := dn.discoveryClient.R().Get(fmt.Sprintf("/users/handle/%s", handle))
	if err != nil {
		return nil, err
	}
	var resp models.UserResponse
	err = resp.UnmarshalBinary(res.Body())
	if err != nil {
		return nil, err
	}

	return resp.Data, nil
}

func (sdk *AudiusSdk) GetUserAiAttributed(handle string, query map[string]string) ([]*models.Track, error) {
	dn := sdk.Discovery
	res, err := dn.discoveryClient.R().SetQueryParams(query).Get(fmt.Sprintf("/users/handle/%s/tracks/ai_attributed", handle))
	if err != nil {
		return nil, err
	}
	var resp models.TracksResponse
	err = resp.UnmarshalBinary(res.Body())
	if err != nil {
		return nil, err
	}

	return resp.Data, nil
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
