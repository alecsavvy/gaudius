package gaudius

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetUploads(t *testing.T) {
	sdk := NewSdkUnsafe()
	uploads, err := sdk.GetUploads()
	require.Nil(t, err)
	require.NotEmpty(t, uploads)
}

func TestGetUpload(t *testing.T) {
	sdk := NewSdkUnsafe()
	id := "01H6EJC9XVMQXM7FA4P0AY148T"
	uploads, err := sdk.GetUpload(id)
	require.Nil(t, err)
	require.NotNil(t, uploads)
	require.EqualValues(t, id, uploads.ID)
}
