package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alecsavvy/gaudius"
	"github.com/davecgh/go-spew/spew"
	"github.com/golang-jwt/jwt"
)

func init() {
	// Register the custom signing method with the jwt package
	jwt.RegisterSigningMethod("keccak256", func() jwt.SigningMethod {
		return &gaudius.Keccak256SigningMethod{}
	})
}

func main() {
	sdk, err := gaudius.NewSdk()
	if err != nil {
		log.Fatal(err)
	}
	port := 8080

	// start webserver to receive auth token
	http.HandleFunc("/oauth", func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.URL.Query().Get("token")
		token, err := jwt.Parse(tokenString, nil)
		if err != nil {
			fmt.Println(err)
		}
		decoded, err := gaudius.ClaimsToDecodedUserToken(token.Claims)
		if err != nil {
			fmt.Println(err)
			return
		}
		spew.Dump(decoded)
	})

	// configure oauth in sdk
	oauthConfig := &gaudius.OauthConfiguration{
		ApiKey:      "",
		Scope:       "write",
		RedirectURI: fmt.Sprintf("http://localhost:%d/oauth", port),
	}
	sdk.ConfigureOauth(oauthConfig)

	// receive oauth url from sdk
	url, err := sdk.OauthUrl()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(url)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
