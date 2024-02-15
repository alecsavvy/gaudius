package gaudius

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTracks(t *testing.T) {
	sdk, err := NewSdk()
	require.Nil(t, err)

	trackID := "W0JP60j"

	track, err := sdk.GetTrack(trackID)
	require.Nil(t, err)
	require.NotNil(t, track)

	trendingTracks, err := sdk.TrendingTracks()
	require.Nil(t, err)
	require.NotEmpty(t, trendingTracks)

	trendingUndergroundTracks, err := sdk.TrendingUndergroundTracks()
	require.Nil(t, err)
	require.NotEmpty(t, trendingUndergroundTracks)
}
