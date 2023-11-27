package gaudius

import (
	"fmt"

	"github.com/alecsavvy/gaudius/gen/discovery/models"
)

func (sdk *AudiusSdk) GetPlaylist(id string) ([]*models.Playlist, error) {
	dn := sdk.Discovery
	res, err := dn.discoveryClient.R().Get(fmt.Sprintf("/playlists/%s", id))
	if err != nil {
		return nil, err
	}
	var resp models.PlaylistResponse
	err = resp.UnmarshalBinary(res.Body())
	if err != nil {
		return nil, err
	}

	return resp.Data, nil
}

func (sdk *AudiusSdk) SearchPlaylists(query string) ([]*models.Playlist, error) {
	dn := sdk.Discovery
	res, err := dn.discoveryClient.R().Get(fmt.Sprintf("/playlists/search?query=%s", query))
	if err != nil {
		return nil, err
	}
	var resp models.PlaylistSearchResult
	err = resp.UnmarshalBinary(res.Body())
	if err != nil {
		return nil, err
	}

	return resp.Data, nil
}

func (sdk *AudiusSdk) TrendingPlaylists() ([]*models.Playlist, error) {
	dn := sdk.Discovery
	res, err := dn.discoveryClient.R().Get("/playlists/trending")
	if err != nil {
		return nil, err
	}
	var resp models.TrendingPlaylistsResponse
	err = resp.UnmarshalBinary(res.Body())
	if err != nil {
		return nil, err
	}

	return resp.Data, nil
}

func (sdk *AudiusSdk) GetPlaylistTracks(id string) ([]*models.Track, error) {
	dn := sdk.Discovery
	res, err := dn.discoveryClient.R().Get(fmt.Sprintf("/playlists/%s/tracks", id))
	if err != nil {
		return nil, err
	}
	var resp models.PlaylistTracksResponse
	err = resp.UnmarshalBinary(res.Body())
	if err != nil {
		return nil, err
	}

	return resp.Data, nil
}
