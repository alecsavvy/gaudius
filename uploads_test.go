package gaudius

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUploads(t *testing.T) {
	sdk := NewSdkUnsafe()
	uploads, err := sdk.GetUploads()
	assert.Nil(t, err)
	assert.NotEmpty(t, uploads)
}

func TestGetUpload(t *testing.T) {
	sdk := NewSdkUnsafe()
	id := "01H6EJC9XVMQXM7FA4P0AY148T"
	uploads, err := sdk.GetUpload(id)
	assert.Nil(t, err)
	assert.NotNil(t, uploads)
	assert.EqualValues(t, id, uploads.ID)
}
