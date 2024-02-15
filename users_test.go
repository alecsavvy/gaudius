package gaudius

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUsers(t *testing.T) {
	sdk, err := NewSdk()
	require.Nil(t, err)

	userID := "1PqKz"
	userHandle := "LemonadeJetpack"

	user, err := sdk.GetUser(userID)
	require.Nil(t, err)
	require.EqualValues(t, userHandle, *user.Handle)

	user, err = sdk.GetUserByHandle(userHandle)
	require.Nil(t, err)
	require.EqualValues(t, userID, *user.ID)

	aiTracks, err := sdk.GetUserAiAttributed(userHandle, nil)
	require.Nil(t, err)
	require.Empty(t, aiTracks)

	users, err := sdk.UserSearch("LemonadeJetpack")
	require.Nil(t, err)
	require.EqualValues(t, userID, *users[0].ID)

	apps, err := sdk.GetUserAuthorizedApps(userID)
	require.Nil(t, err)
	require.NotEmpty(t, apps)

	tracks, err := sdk.GetUserTracks(*users[0].ID, map[string]string{"limit": "10", "filter_tracks": "all"})
	require.Nil(t, err)
	require.NotEmpty(t, tracks)
}

func TestFullUsers(t *testing.T) {
	sdk, err := NewSdk()
	require.Nil(t, err)

	topPopUsers, err := sdk.GetTopGenreUsers(Genre.Pop, nil)
	require.Nil(t, err)
	require.NotEmpty(t, topPopUsers)
}
