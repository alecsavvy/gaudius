package gaudius

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestImageDimensions(t *testing.T) {
	small := ImageDimensions("sm")
	require.EqualValues(t, "150x150", small)
	small = ImageDimensions("small")
	require.EqualValues(t, "150x150", small)

	medium := ImageDimensions("md")
	require.EqualValues(t, "400x400", medium)
	medium = ImageDimensions("medium")
	require.EqualValues(t, "400x400", medium)

	large := ImageDimensions("lg")
	require.EqualValues(t, "1000x1000", large)
	large = ImageDimensions("large")
	require.EqualValues(t, "1000x1000", large)

	large = ImageDimensions("")
	require.EqualValues(t, "1000x1000", large)
}
