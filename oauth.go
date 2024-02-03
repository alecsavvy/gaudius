package gaudius

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/alecsavvy/gaudius/gen/discovery/models"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/sha3"
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

// this is because the swagger spec has a string for a number iat and fails to serialize
// also, claims aren't marshallable
func ClaimsToDecodedUserToken(claims jwt.Claims) (*models.DecodedUserToken, error) {
	claimsMap, ok := claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("could not convert claims to map claims")
	}

	decodedToken := &models.DecodedUserToken{}

	// Extract and assert the type of each field from claimsMap
	if email, ok := claimsMap["email"].(string); ok {
		decodedToken.Email = &email
	}

	if handle, ok := claimsMap["handle"].(string); ok {
		decodedToken.Handle = &handle
	}

	if iat, ok := claimsMap["iat"].(float64); ok {
		iatStr := strconv.FormatFloat(iat, 'f', 0, 64)
		decodedToken.Iat = &iatStr
	}

	if name, ok := claimsMap["name"].(string); ok {
		decodedToken.Name = &name
	}

	if profilePicture, ok := claimsMap["profilePicture"].(map[string]interface{}); ok {
		profilePictureJSON, err := json.Marshal(profilePicture)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Error marshaling profile picture: %v", err))
		}
		profilePic := &models.ProfilePicture{}
		err = json.Unmarshal(profilePictureJSON, profilePic)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Error unmarshaling profile picture into struct: %v", err))
		}
		decodedToken.ProfilePicture = profilePic
	}

	if sub, ok := claimsMap["sub"].(string); ok {
		decodedToken.Sub = &sub
	}

	if userID, ok := claimsMap["userId"].(string); ok {
		decodedToken.UserID = &userID
	}

	if verified, ok := claimsMap["verified"].(bool); ok {
		decodedToken.Verified = &verified
	}

	return decodedToken, nil
}

type Keccak256SigningMethod struct{}

func (m *Keccak256SigningMethod) Alg() string {
	return "keccak256"
}

func (m *Keccak256SigningMethod) Sign(signingString string, key interface{}) (string, error) {
	// Convert the key to []byte if it's not already
	var keyBytes []byte
	switch k := key.(type) {
	case []byte:
		keyBytes = k
	case string:
		keyBytes = []byte(k)
	default:
		return "", jwt.ErrInvalidKeyType
	}

	// Create a Keccak256 hasher
	hasher := sha3.NewLegacyKeccak256()
	// Write the signingString to the hasher; this is the data you're signing
	_, err := hasher.Write([]byte(signingString))
	if err != nil {
		return "", err
	}
	// Sum the hash, appending it to the keyBytes (if your signing logic requires the key in the hash)
	signature := hasher.Sum(keyBytes)
	// Convert the binary signature to a hex string for easier handling
	signatureString := hex.EncodeToString(signature)
	return signatureString, nil
}

func (m *Keccak256SigningMethod) Verify(signingString, signature string, key interface{}) error {
	// Recreate the signature using the signingString and key, then compare it to the provided signature
	newSig, err := m.Sign(signingString, key)
	if err != nil {
		return err
	}

	// Check if the new signature matches the provided signature
	if hex.EncodeToString([]byte(newSig)) != signature {
		return jwt.ErrSignatureInvalid
	}
	return nil
}
