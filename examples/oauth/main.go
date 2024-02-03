package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/alecsavvy/gaudius"
	"github.com/davecgh/go-spew/spew"
	"github.com/golang-jwt/jwt"
)

func init() {
	// Register the custom signing method with the jwt package
	jwt.RegisterSigningMethod("Keccak256", func() jwt.SigningMethod {
		return &gaudius.Keccak256SigningMethod{}
	})
}

func main() {
	sdk := gaudius.NewSdkUnsafe()
	port := 8080

	// start webserver to receive auth token
	http.HandleFunc("/oauth", func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.URL.Query().Get("token")
		token, _ := jwt.Parse(tokenString, nil)
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

func base64UrlDecode(data string) ([]byte, error) {
	// Add padding characters if necessary
	if m := len(data) % 4; m != 0 {
		data += strings.Repeat("=", 4-m)
	}

	// Base64 URL decode
	bytes, err := base64.URLEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}
