package main

import (
	"fmt"
	"log"

	"github.com/alecsavvy/gaudius"
)

func main() {
	sdk := gaudius.NewSdkUnsafe()
	oauthConfig := &gaudius.OauthConfiguration{
		ApiKey:      "",
		Scope:       "write",
		RedirectURI: "http://localhost:8080/oauth",
	}
	sdk.ConfigureOauth(oauthConfig)

	url, err := sdk.OauthUrl()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(url)
}
