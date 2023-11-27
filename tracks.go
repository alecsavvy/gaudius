package gaudius

import (
	"fmt"

	"github.com/alecsavvy/gaudius/gen/discovery/models"
)

func (sdk *AudiusSdk) GetTrack(id string) (*models.Track, error) {
	dn := sdk.Discovery
	res, err := dn.discoveryClient.R().Get(fmt.Sprintf("/tracks/%s", id))
	if err != nil {
		return nil, err
	}
	var resp models.TrackResponse
	err = resp.UnmarshalBinary(res.Body())
	if err != nil {
		return nil, err
	}

	return resp.Data, nil
}

func (sdk *AudiusSdk) SearchTracks(query string) ([]*models.Track, error) {
	dn := sdk.Discovery
	res, err := dn.discoveryClient.R().Get(fmt.Sprintf("/tracks/search?query=%s", query))
	if err != nil {
		return nil, err
	}
	var resp models.TrackSearch
	err = resp.UnmarshalBinary(res.Body())
	if err != nil {
		return nil, err
	}

	return resp.Data, nil
}

func (sdk *AudiusSdk) TrendingTracks() ([]*models.Track, error) {
	dn := sdk.Discovery
	res, err := dn.discoveryClient.R().Get("/tracks/trending")
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

func (sdk *AudiusSdk) TrendingUndergroundTracks() ([]*models.Track, error) {
	dn := sdk.Discovery
	res, err := dn.discoveryClient.R().Get("/tracks/trending/underground")
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
