package gaudius

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetImage(t *testing.T) {
	sdk := NewSdkUnsafe()
	
	image, err := sdk.GetImage("01H6EJC9XVMQXM7FA4P0AY148T")
	assert.Nil(t, err)
	assert.NotEmpty(t, image)
	
	expected, err := ioutil.ReadFile("test_assets/test_image_large.jpg")
	assert.Nil(t, err)
	assert.NotEmpty(t, expected)

	assert.EqualValues(t, expected, image)
}