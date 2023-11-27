package gaudius

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetImage(t *testing.T) {
	sdk := NewSdkUnsafe()

	image, err := sdk.GetImage("01H6EJC9XVMQXM7FA4P0AY148T")
	require.Nil(t, err)
	require.NotEmpty(t, image)

	expected, err := ioutil.ReadFile("test_assets/test_image_large.jpg")
	require.Nil(t, err)
	require.NotEmpty(t, expected)

	require.EqualValues(t, expected, image)
}
