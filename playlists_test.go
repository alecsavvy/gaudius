package gaudius

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPlaylists(t *testing.T) {
	sdk, err := NewSdk()
	require.Nil(t, err)

	playlistID := "LjErB"

	playlists, err := sdk.GetPlaylist(playlistID)
	require.Nil(t, err)
	require.NotEmpty(t, playlists)

	tracks, err := sdk.GetPlaylistTracks(playlistID)
	require.Nil(t, err)
	require.NotEmpty(t, tracks)

	trendingplaylists, err := sdk.TrendingPlaylists()
	require.Nil(t, err)
	require.NotEmpty(t, trendingplaylists)
}
