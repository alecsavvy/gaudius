package gaudius

import (
	"errors"
	"fmt"
)

type OauthConfiguration struct {
	Scope        string
	ApiKey       string
	RedirectURI  string
	AppOrigin    string
	State        string
	ResponseMode string
}

func (sdk *AudiusSdk) ConfigureOauth(config *OauthConfiguration) {
	sdk.Oauth = config
}

func (sdk *AudiusSdk) OauthUrl() (string, error) {
	config := sdk.Oauth
	if config.Scope == "" {
		config.Scope = "read"
	}
	scope := fmt.Sprintf("scope=%s", config.Scope)
	if config.ApiKey == "" {
		return "", errors.New("ApiKey is a required field for oauth configuration")
	}
	apiKey := fmt.Sprintf("api_key=%s", config.ApiKey)
	if config.RedirectURI == "" {
		return "", errors.New("RedirectURI is a required field for oauth configuration")
	}
	redirectUri := fmt.Sprintf("redirect_uri=%s", config.RedirectURI)
	if config.ResponseMode == "" {
		// default to query since servers can't read fragment fields
		config.ResponseMode = "query"
	}
	responseMode := fmt.Sprintf("response_mode=%s", config.ResponseMode)

	return fmt.Sprintf("https://audius.co/oauth/auth?%s&%s&%s&%s", scope, apiKey, redirectUri, responseMode), nil
}
