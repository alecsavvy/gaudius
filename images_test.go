package gaudius

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetImage(t *testing.T) {
	sdk := NewSdk()
	image, err := sdk.GetImage("01H6EJC9XVMQXM7FA4P0AY148T")
	assert.Nil(t, err)
	assert.NotEmpty(t, image)
}