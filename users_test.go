package gaudius

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUsers(t *testing.T) {
	sdk := NewSdkUnsafe()

	userID := "1PqKz"
	userHandle := "LemonadeJetpack"

	user, err := sdk.GetUser(userID)
	require.Nil(t, err)
	require.EqualValues(t, userHandle, *user.Handle)

	user, err = sdk.GetUserByHandle(userHandle)
	require.Nil(t, err)
	require.EqualValues(t, userID, *user.ID)

	tracks, err := sdk.GetUserAiAttributed(userHandle)
	require.Nil(t, err)
	require.Empty(t, tracks)

	users, err := sdk.UserSearch("LemonadeJetpack")
	require.Nil(t, err)
	require.EqualValues(t, userID, *users[0].ID)

	apps, err := sdk.GetUserAuthorizedApps(userID)
	require.Nil(t, err)
	require.NotEmpty(t, apps)

}
