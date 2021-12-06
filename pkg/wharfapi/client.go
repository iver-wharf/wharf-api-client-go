package wharfapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// AuthError is returned on authentication/authorization errors issued when
// trying to communicate with the Wharf API.
//
// This could be because of missing, invalid, or outdated authentication header
// provided to the client.
type AuthError struct {
	Realm string
}

func (e *AuthError) Error() string {
	return e.Realm
}

// Client contains authentication and API URLs used to access
// the Wharf main API.
type Client struct {
	AuthHeader string
	APIURL     string
}

// WharfClient contains authentication and API URLs used to access
// the Wharf main API.
//
// Deprecated: This type has been renamed to Client and may be removed in a
// future release.
type WharfClient Client

func (c Client) Get(from, path string, q url.Values) ([]byte, error) {
	return doRequestNew(fmt.Sprintf("GET | %s", from), http.MethodGet, c.APIURL, path, q, nil, c.AuthHeader)
}

func (c Client) GetDecoded(response interface{}, from, path string, q url.Values) error {
	bytes, err := c.Get(from, path, q)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, response)
}

func (c Client) Post(from, path string, q url.Values, body []byte) ([]byte, error) {
	return doRequestNew(fmt.Sprintf("POST | %s", from), http.MethodPost, c.APIURL, path, q, body, c.AuthHeader)
}

func (c Client) PostDecoded(response interface{}, from, path string, q url.Values, body []byte) error {
	bytes, err := c.Post(from, path, q, body)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, response)
}

func (c Client) Put(from, path string, q url.Values, body []byte) ([]byte, error) {
	return doRequestNew(fmt.Sprintf("PUT | %s", from), http.MethodGet, c.APIURL, path, q, body, c.AuthHeader)
}
func (c Client) PutDecoded(response interface{}, from, path string, q url.Values, body []byte) error {
	bytes, err := c.Put(from, path, q, body)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, response)
}
