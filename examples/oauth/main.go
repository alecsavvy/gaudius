package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/alecsavvy/gaudius"
	"github.com/alecsavvy/gaudius/gen/discovery/models"
)

func main() {
	sdk := gaudius.NewSdkUnsafe()
	port := 8080

	// start webserver to receive auth token
	http.HandleFunc("/oauth", func(w http.ResponseWriter, r *http.Request) {
		token := r.URL.Query().Get("token")

		fmt.Println(token)

		claims := strings.Split(token, ".")[1]
		bytes, err := base64UrlDecode(claims)
		if err != nil {
			fmt.Printf("Error decoding payload: %v\n", err)
			return
		}

		decodedToken := &models.DecodedUserToken{}
		err = decodedToken.UnmarshalBinary(bytes)
		if err != nil {
			fmt.Printf("Error unmarshaling payload: %v\n", err)
			return
		}

		fmt.Println(decodedToken)
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
