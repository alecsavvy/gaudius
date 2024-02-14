package gaudius

import (
	"fmt"
	"math/rand"
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

// for parity with the web3.js web3.utils.utf8ToHex() call
func Utf8ToHex(s string) [32]byte {
	hex := [32]byte{}
	copy(hex[:], s)
	return hex
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

func shuffle[T any](slice []T) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
}
