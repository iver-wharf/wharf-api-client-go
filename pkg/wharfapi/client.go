package wharfapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/blang/semver/v4"
	"github.com/iver-wharf/wharf-core/pkg/logger"
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

var (
	// ErrOutdatedServer is returned from an endpoint method when the
	// Client.ErrIfOutdatedServer flag is enabled and the server is of a lower
	// version than when the endpoint was first introduced to the wharf-api.
	ErrOutdatedServer = errors.New("outdated server")
	// ErrOutdatedClient is returned from an endpoint method when the
	// Client.ErrIfOutdatedClient flag is enabled and the client is of a too
	// low version than when the server.
	ErrOutdatedClient = errors.New("outdated client")
)

// Client contains authentication and API URLs used to access
// the Wharf main API.
type Client struct {
	AuthHeader string
	APIURL     string

	// ErrIfOutdatedClient will error if the client is outdated. Wharf aims
	// for a backward compatability of 1 major version back, so the client will
	// only prematurely error before making a request if the client is 2 major
	// versions behind or more. Example:
	//
	//   Server version        Client supports        Client outdated?
	//   v5.0.0                v5.0.0                 No
	//   v5.1.0                v5.0.0                 No
	//   v6.0.0                v5.0.0                 No
	//   v6.12.5               v5.0.0                 No
	//   v7.0.0                v5.0.0                 Yes
	ErrIfOutdatedClient bool

	// ErrIfOutdatedServer will error if the remote API version is too low for
	// each endpoint before even making the web request.
	ErrIfOutdatedServer bool

	// DisableOutdatedLogging will disable the logging to console if there are
	// unattended issues regarding version mismatch between the client and the
	// server.
	DisableOutdatedLogging bool

	hasCheckedVersion             bool
	hasLoggedClientVersionWarning bool
	hasLoggedServerVersionWarning bool
	cachedVersion                 *semver.Version
}

// HighestSupportedVersion is the highest version that the wharf-api-client-go
// is known to work for. It is used when checking if the client is outdated,
// given the Client.ErrIfOutdatedClient is enabled.
var HighestSupportedVersion = semver.MustParse("5.1.0")

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

// SetCachedVersion will override the version that the wharf-api-client-go
// thinks the remote API has when validating the Client.ErrIfOutdatedServer.
func (c *Client) SetCachedVersion(major, minor, patch uint64) {
	c.cachedVersion = &semver.Version{
		Major: major,
		Minor: minor,
		Patch: patch,
	}
	c.hasCheckedVersion = true
	c.hasLoggedClientVersionWarning = false
}

// ResetCachedVersion will reset the version that the wharf-api client thinks
// the remote API has, and will then check for a fresh value on the next request.
func (c *Client) ResetCachedVersion() {
	c.cachedVersion = nil
	c.hasCheckedVersion = false
	c.hasLoggedClientVersionWarning = false
}

func (c *Client) getCachedOrFetchedVersion() *semver.Version {
	if c.hasCheckedVersion {
		return c.cachedVersion
	}
	_, err := c.GetVersion()
	if err != nil || c.cachedVersion == nil {
		return nil
	}
	log.Debug().
		WithStringer("version", c.cachedVersion).
		Message("Detected server version.")
	return c.cachedVersion
}

func (c *Client) validateEndpointVersion(major, minor, patch uint64) error {
	if !c.ErrIfOutdatedClient && !c.ErrIfOutdatedServer && c.DisableOutdatedLogging {
		// micro-optimization:
		// skip fetching version if we don't even care about versions
		return nil
	}
	apiVersion := c.getCachedOrFetchedVersion()
	return c.validateEndpointVersionNoLookup(major, minor, patch, apiVersion)
}

func (c *Client) validateEndpointVersionNoLookup(major, minor, patch uint64, apiVersion *semver.Version) error {
	if apiVersion == nil {
		return nil
	}
	endpointVersion := semver.Version{
		Major: major,
		Minor: minor,
		Patch: patch,
	}
	if err := c.validateServerVersion(*apiVersion, endpointVersion); err != nil {
		if !c.DisableOutdatedLogging && !c.hasLoggedServerVersionWarning {
			c.hasLoggedServerVersionWarning = true
			log.Warn().WithError(err).
				Message("Server is outdated.")
		}
		if c.ErrIfOutdatedServer {
			return err
		}
	}
	if level, err := c.validateClientVersion(*apiVersion); err != nil {
		if !c.DisableOutdatedLogging && !c.hasLoggedClientVersionWarning {
			c.hasLoggedClientVersionWarning = true
			logger.NewEventFromLogger(log, level).
				WithError(err).
				Message("Client is outdated.")
		}
		if c.ErrIfOutdatedClient && level >= logger.LevelError {
			return err
		}
	}
	return nil
}

func (c *Client) validateServerVersion(apiVersion, endpointVersion semver.Version) error {
	if apiVersion.LT(endpointVersion) {
		return fmt.Errorf("%w: %s (server) is less than %s (when endpoint was introduced)",
			ErrOutdatedServer, apiVersion, endpointVersion)
	}
	return nil
}

func (c *Client) validateClientVersion(apiVersion semver.Version) (logger.Level, error) {
	if HighestSupportedVersion.Major+1 < apiVersion.Major {
		return logger.LevelError, fmt.Errorf(
			"%w: %s (server) is too new for %s (highest supported version by client)",
			ErrOutdatedClient, apiVersion, HighestSupportedVersion)
	}
	if HighestSupportedVersion.Major < apiVersion.Major {
		return logger.LevelWarn, fmt.Errorf(
			"%w: %s (server) is newer than %s (highest supported version by"+
				" client), enough to still be supported, but use with caution",
			ErrOutdatedClient, apiVersion, HighestSupportedVersion)
	} else if HighestSupportedVersion.LT(apiVersion) {
		return logger.LevelDebug, fmt.Errorf(
			"%w: %s (server) is slightly newer than %s (highest supported "+
				"version by client), however way within the supported range",
			ErrOutdatedClient, apiVersion, HighestSupportedVersion)
	}
	return 0, nil
}
