package main

import (
	"bytes"
	"fmt"
	"image"
	"log"

	"github.com/alecsavvy/gaudius"
	"github.com/kr/pretty"
	"github.com/qeesung/image2ascii/convert"
)

func main() {
	sdk := gaudius.NewSdk()
	
	user, err := sdk.GetUser("1PqKz")

	if err != nil {
		log.Fatal(err)
	}

	// print user to console
	fmt.Printf("%# v", pretty.Formatter(user))

	imgdata, err := sdk.GetImage("01H6EJC9XVMQXM7FA4P0AY148T")

	if err != nil {
		log.Fatal(err)
	}

	// print image to console
	img, _, err := image.Decode(bytes.NewReader(imgdata))
	converter := convert.NewImageConverter()
	opts := convert.DefaultOptions
	fmt.Println(converter.Image2ASCIIString(img, &opts))
}