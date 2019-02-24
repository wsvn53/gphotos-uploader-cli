package tokenstore

import (
	"encoding/json"
	"fmt"
	//"github.com/zalando/go-keyring"
	"golang.org/x/oauth2"
	"io/ioutil"
)

const (
	serviceName = "googlephotos-uploader-go-api"
)

var (
	// ErrNotFound is the expected error if the token isn't found in the keyring
	ErrNotFound = fmt.Errorf("failed retrieving token from keyring")

	// ErrInvalidToken is the expected error if the token isn't a valid one
	ErrInvalidToken = fmt.Errorf("invalid token")
)

// StoreToken lets you store a token in the OS keyring
func StoreToken(googleUserEmail string, token *oauth2.Token) error {
	tokenJSONBytes, err := json.Marshal(token)
	if err != nil {
		return err
	}

	storeFile := fmt.Sprintf("./oauth.store_%s_%s", serviceName, googleUserEmail)
	err = ioutil.WriteFile(storeFile, tokenJSONBytes, 0777)

	//err = keyring.Set(serviceName, googleUserEmail, string(tokenJSONBytes))
	if err != nil {
		return fmt.Errorf("failed storing token into keyring: %v", err)
	}
	return nil
}

// RetrieveToken lets you get a token by google account email
func RetrieveToken(googleUserEmail string) (*oauth2.Token, error) {
	//tokenJSONString, err := keyring.Get(serviceName, googleUserEmail)
	storeFile := fmt.Sprintf("./oauth.store_%s_%s", serviceName, googleUserEmail)
	tokenJSONBytes, err := ioutil.ReadFile(storeFile)
	tokenJSONString := string(tokenJSONBytes)
	if err != nil {
		//if err == keyring.ErrNotFound {
		//	return nil, ErrNotFound
		//}
		return nil, err
	}

	var token oauth2.Token
	err = json.Unmarshal([]byte(tokenJSONString), &token)
	if err != nil {
		return nil, fmt.Errorf("failed unmarshaling token: %v", err)
	}

	// validate token
	if !token.Valid() {
		return nil, ErrInvalidToken
	}

	return &token, nil
}

// MockInit sets the provider to a mocked memory store, using keyring mock
func MockInit() {
	//keyring.MockInit()
}
