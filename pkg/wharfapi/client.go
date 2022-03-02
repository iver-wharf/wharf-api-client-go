package wharfapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/Masterminds/semver/v3"
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

	// ErrIfOutdatedClient will error if the client is outdated. Wharf aims
	// for a backward compatability of 1 major version back, so the client will
	// only prematurely error before making a request if the client is 2 major
	// versions behind or more. Example:
	//   Server version        Client supports        Client outdated?
	//   v5.0.0                v5.0.0                 No
	//   v6.0.0                v5.0.0                 No
	//   v6.12.5               v5.0.0                 No
	//   v7.0.0                v5.0.0                 Yes
	ErrIfOutdatedClient bool

	// ErrIfOutdatedServer will error if the remote API version is too low for
	// each endpoint before even making the web request.
	ErrIfOutdatedServer bool

	hasCheckedVersion bool
	cachedVersion     *semver.Version
}

// HighestSupportedVersion is the highest version that the wharf-api-client-go
// is known to work for.
var HighestSupportedVersion = semver.MustParse("v5.1.0")

// WharfClient contains authentication and API URLs used to access
// the Wharf main API.
//
// Deprecated: This type has been renamed to Client and may be removed in a
// future release.
type WharfClient Client

func (c *Client) get(path string, q url.Values) (io.ReadCloser, error) {
	req, err := c.newRequest(http.MethodGet, path, q, nil)
	if err != nil {
		return nil, err
	}
	return doRequest(req)
}

func (c *Client) getUnmarshal(path string, q url.Values, response interface{}) error {
	ioBody, err := c.get(path, q)
	if err != nil {
		return err
	}
	err = json.NewDecoder(ioBody).Decode(response)
	if err != nil {
		return err
	}
	return ioBody.Close()
}

func (c *Client) post(path string, q url.Values, body []byte) (io.ReadCloser, error) {
	req, err := c.newRequest(http.MethodPost, path, q, body)
	if err != nil {
		return nil, err
	}
	return doRequest(req)
}

func (c *Client) postJSON(path string, q url.Values, obj interface{}) (io.ReadCloser, error) {
	bodyBytes, err := json.Marshal(&obj)
	if err != nil {
		return nil, err
	}
	return c.post(path, q, bodyBytes)
}

func (c *Client) postJSONUnmarshal(path string, q url.Values, obj interface{}, response interface{}) error {
	ioBody, err := c.postJSON(path, q, obj)
	if err != nil {
		return err
	}
	err = json.NewDecoder(ioBody).Decode(response)
	if err != nil {
		return err
	}
	return ioBody.Close()
}

func (c *Client) put(path string, q url.Values, body []byte) (io.ReadCloser, error) {
	req, err := c.newRequest(http.MethodPut, path, q, body)
	if err != nil {
		return nil, err
	}
	return doRequest(req)
}

func (c *Client) putJSON(path string, q url.Values, obj interface{}) (io.ReadCloser, error) {
	bodyBytes, err := json.Marshal(&obj)
	if err != nil {
		return nil, err
	}
	return c.put(path, q, bodyBytes)
}

func (c *Client) putJSONUnmarshal(path string, q url.Values, obj interface{}, response interface{}) error {
	ioBody, err := c.putJSON(path, q, obj)
	if err != nil {
		return err
	}
	err = json.NewDecoder(ioBody).Decode(response)
	if err != nil {
		return err
	}
	return ioBody.Close()
}

func (c *Client) newRequest(method, path string, q url.Values, body []byte) (*http.Request, error) {
	return newRequest(method, c.AuthHeader, c.APIURL, path, q, body)
}

func (c *Client) getCachedOrFetchedVersion() *semver.Version {
	if c.hasCheckedVersion {
		return c.cachedVersion
	}
	c.hasCheckedVersion = true
	resp, err := c.GetVersion()
	if err != nil {
		log.Warn().WithError(err).
			Message("Failed to get version from API. Version validation is turned off. Use at your own risk.")
		return nil
	}
	v, err := semver.NewVersion(resp.Version)
	if err != nil {
		log.Warn().WithError(err).
			WithString("version", resp.Version).
			Message("Failed to parse version from API. Version validation is turned off. Use at your own risk.")
		return nil
	}
	log.Debug().
		WithStringer("version", v).
		Message("Detected server version.")
	c.cachedVersion = v
	return v
}

func (c *Client) validateEndpointVersion(endpointVersion *semver.Version) error {
	if !c.ErrIfOutdatedClient && !c.ErrIfOutdatedServer {
		return nil
	}
	apiVersion := c.getCachedOrFetchedVersion()
	if apiVersion == nil {
		return nil
	}
	if c.ErrIfOutdatedServer && apiVersion.LessThan(endpointVersion) {
		return fmt.Errorf("%w: %s (server) is less than %s (when endpoint was introduced)",
			ErrOutdatedServer, apiVersion, endpointVersion)
	}
	if c.ErrIfOutdatedClient && HighestSupportedVersion.Major()+1 < apiVersion.Major() {
		return fmt.Errorf("%w: %s (server) is too new for %s (highest supported version by client)",
			ErrOutdatedClient, apiVersion, HighestSupportedVersion)
	}
	return nil
}
