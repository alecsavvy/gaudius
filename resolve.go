package gaudius

import (
	"fmt"
)

func (sdk *AudiusSdk) Resolve(url string) (*string, error) {
	dn := sdk.Discovery
	res, err := dn.discoveryClient.R().Get(fmt.Sprintf("/resolve/?url=%s", url))
	if err != nil {
		return nil, err
	}

	// this endpoint redirects, extract the url
	redirect, err := res.RawResponse.Location()
	if err != nil {
		return nil, err
	}

	redirectstr := redirect.String()

	return &redirectstr, nil
}
