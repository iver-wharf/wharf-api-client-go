package wharfapi

import (
	"encoding/json"
	"io"
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

func (c Client) get(path string, q url.Values) (*io.ReadCloser, error) {
	req, err := c.newRequest(http.MethodGet, path, q, nil)
	if err != nil {
		return nil, err
	}
	return doRequest(req)
}

func (c Client) getUnmarshal(path string, q url.Values, response interface{}) error {
	ioBody, err := c.get(path, q)
	if err != nil {
		return err
	}
	return json.NewDecoder(*ioBody).Decode(response)
}

func (c Client) post(path string, q url.Values, body []byte) (*io.ReadCloser, error) {
	req, err := c.newRequest(http.MethodPost, path, q, body)
	if err != nil {
		return nil, err
	}
	return doRequest(req)
}

func (c Client) postJSON(path string, q url.Values, obj interface{}) (*io.ReadCloser, error) {
	bodyBytes, err := json.Marshal(&obj)
	if err != nil {
		return nil, err
	}
	return c.post(path, q, bodyBytes)
}

func (c Client) postJSONUnmarshal(path string, q url.Values, obj interface{}, response interface{}) error {
	ioBody, err := c.postJSON(path, q, obj)
	if err != nil {
		return err
	}
	return json.NewDecoder(*ioBody).Decode(response)
}

func (c Client) put(path string, q url.Values, body []byte) (*io.ReadCloser, error) {
	req, err := c.newRequest(http.MethodPut, path, q, body)
	if err != nil {
		return nil, err
	}
	return doRequest(req)
}

func (c Client) putJSON(path string, q url.Values, obj interface{}) (*io.ReadCloser, error) {
	bodyBytes, err := json.Marshal(&obj)
	if err != nil {
		return nil, err
	}
	return c.put(path, q, bodyBytes)
}

func (c Client) putJSONUnmarshal(path string, q url.Values, obj interface{}, response interface{}) error {
	ioBody, err := c.putJSON(path, q, obj)
	if err != nil {
		return err
	}
	return json.NewDecoder(*ioBody).Decode(response)
}

func (c Client) newRequest(method, path string, q url.Values, body []byte) (*http.Request, error) {
	return newRequest(method, c.AuthHeader, c.APIURL, path, q, body)
}
