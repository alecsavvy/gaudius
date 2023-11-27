package gaudius

import (
	"fmt"

	"github.com/alecsavvy/gaudius/gen/discovery/models"
)

type GetTipsParams struct {
	offset *bool
	limit *uint64
	user_id *string
	receiver_min_followers *uint64
	receiver_is_verified *bool
	current_user_follows *string
	unique_by *string
}

func (g *GetTipsParams) IntoQueryString() string {
    var queryString string

    if g.offset != nil {
        queryString += fmt.Sprintf("offset=%t&", *g.offset)
    }
    if g.limit != nil {
        queryString += fmt.Sprintf("limit=%d&", *g.limit)
    }
    if g.user_id != nil {
        queryString += fmt.Sprintf("user_id=%s&", *g.user_id)
    }
    if g.receiver_min_followers != nil {
        queryString += fmt.Sprintf("receiver_min_followers=%d&", *g.receiver_min_followers)
    }
    if g.receiver_is_verified != nil {
        queryString += fmt.Sprintf("receiver_is_verified=%t&", *g.receiver_is_verified)
    }
    if g.current_user_follows != nil {
        queryString += fmt.Sprintf("current_user_follows=%s&", *g.current_user_follows)
    }
    if g.unique_by != nil {
        queryString += fmt.Sprintf("unique_by=%s&", *g.unique_by)
    }

    // Remove the trailing '&' if it exists
    if len(queryString) > 0 {
        queryString = queryString[:len(queryString)-1]
    }

    return queryString
}

func (sdk *AudiusSdk) GetTips(params GetTipsParams) (*models.User, error) {
	dn := sdk.Discovery
	qs := params.IntoQueryString()
	res, err := dn.discoveryClient.R().Get(fmt.Sprintf("/tips?%s", qs))
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