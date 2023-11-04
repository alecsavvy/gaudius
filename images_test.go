package gaudius

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetImage(t *testing.T) {
	sdk := NewSdk()
	_, err := sdk.GetImage("01H6EJC9XVMQXM7FA4P0AY148T")
	assert.Nil(t, err)
	
}