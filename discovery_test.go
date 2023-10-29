package gaudius

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	sdk := NewSdk()
	
	user, err := sdk.getUser("1PqKz")
	assert.Nil(t, err)

	assert.EqualValues(t, "LemonadeJetpack", *user.Handle)
}


