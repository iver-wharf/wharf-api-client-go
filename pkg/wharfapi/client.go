package wharfapi

import (
	"encoding/json"
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

func (c Client) Get(path string, q url.Values) ([]byte, error) {
	req, err := c.NewRequest(http.MethodGet, path, q, nil)
	if err != nil {
		return nil, err
	}
	return doRequest(req)
}

func (c Client) GetUnmarshal(path string, q url.Values, response interface{}) error {
	bytes, err := c.Get(path, q)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, response)
}

func (c Client) Post(path string, q url.Values, body []byte) ([]byte, error) {
	req, err := c.NewRequest(http.MethodPost, path, q, body)
	if err != nil {
		return nil, err
	}
	return doRequest(req)
}

func (c Client) PostJSON(path string, q url.Values, obj interface{}) ([]byte, error) {
	bodyBytes, err := json.Marshal(&obj)
	if err != nil {
		return nil, err
	}
	return c.Post(path, q, bodyBytes)
}

func (c Client) PostJSONUnmarshal(path string, q url.Values, obj interface{}, response interface{}) error {
	responseBytes, err := c.PostJSON(path, q, obj)
	if err != nil {
		return err
	}
	return json.Unmarshal(responseBytes, response)
}

func (c Client) Put(path string, q url.Values, body []byte) ([]byte, error) {
	req, err := c.NewRequest(http.MethodPut, path, q, body)
	if err != nil {
		return nil, err
	}
	return doRequest(req)
}

func (c Client) PutJSON(path string, q url.Values, obj interface{}) ([]byte, error) {
	bodyBytes, err := json.Marshal(&obj)
	if err != nil {
		return nil, err
	}
	return c.Put(path, q, bodyBytes)
}

func (c Client) PutJSONUnmarshal(path string, q url.Values, obj interface{}, response interface{}) error {
	responseBytes, err := c.PutJSON(path, q, obj)
	if err != nil {
		return err
	}
	return json.Unmarshal(responseBytes, response)
}

func (c Client) NewRequest(method, path string, q url.Values, body []byte) (*http.Request, error) {
	return newRequest(method, c.AuthHeader, c.APIURL, path, q, body)
}
