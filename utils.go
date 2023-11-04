package gaudius

import (
	"fmt"
)

// formats to "dimensionXdimension" for content nodes
func ImageSizeFormatter(dimension int) string {
	return fmt.Sprintf("%dx%d", dimension, dimension)
}

// takes in english dimensions like "small"/"large" and converts
// to common content node numeric dimensions
// defaults to large image should size be invalid
func ImageDimensions(size string) string {
	dimensions := 1000
	switch size {
	case "small", "sm":
		dimensions = 150
	case "medium", "md":
		dimensions = 400
	}
	return ImageSizeFormatter(dimensions)
}