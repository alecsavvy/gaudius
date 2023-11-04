package gaudius

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImageDimensions(t *testing.T) {
	small := ImageDimensions("sm")
	assert.EqualValues(t, "150x150", small)
	small = ImageDimensions("small")
	assert.EqualValues(t, "150x150", small)

	medium := ImageDimensions("md")
	assert.EqualValues(t, "400x400", medium)
	medium = ImageDimensions("medium")
	assert.EqualValues(t, "400x400", medium)

	large := ImageDimensions("lg")
	assert.EqualValues(t, "1000x1000", large)
	large = ImageDimensions("large")
	assert.EqualValues(t, "1000x1000", large)
	
	large = ImageDimensions("")
	assert.EqualValues(t, "1000x1000", large)
}