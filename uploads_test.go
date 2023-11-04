package gaudius

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUploads(t *testing.T) {
	sdk := NewSdk()
	uploads, err := sdk.GetUploads()
	assert.Nil(t, err)
	assert.NotEmpty(t, uploads)
}