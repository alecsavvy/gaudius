package gaudius

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUsers(t *testing.T) {
	sdk := NewSdk()
	
	userID := "1PqKz"
	userHandle := "LemonadeJetpack"

	user, err := sdk.GetUser(userID)
	assert.Nil(t, err)
	assert.EqualValues(t, userHandle, *user.Handle)

	user, err = sdk.GetUserHandle(userHandle)
	assert.Nil(t, err)
	assert.EqualValues(t, userID, *user.ID)

	tracks, err := sdk.GetUserAiAttributed(userHandle)
	assert.Nil(t, err)
	assert.Empty(t, tracks)

	users, err := sdk.UserSearch("LemonadeJetpack")
	assert.Nil(t, err)
	assert.EqualValues(t, userID, *users[0].ID)
}
