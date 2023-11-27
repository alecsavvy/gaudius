package gaudius

import (
	"fmt"
	"time"
)

/** Exported utilities */

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

/** Internal Utilities */

// asynchronously awaits for a period of time
// use like: <-await(time.Second)
// for awaiting one second
func await(d time.Duration) <-chan bool {
	c := make(chan bool)

	go func() {
		time.Sleep(d)
		c <- true
		close(c)
	}()

	return c
}
