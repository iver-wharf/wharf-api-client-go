package wharfapi

import (
	"github.com/blang/semver/v4"
	"github.com/iver-wharf/wharf-core/pkg/app"
)

// GetVersion gets the version of the API by invoking the
// HTTP request:
//  GET /api/version
//
// Added in wharf-api v4.0.0.
func (c *Client) GetVersion() (app.Version, error) {
	if err := c.validateEndpointVersionNoLookup(4, 0, 0, c.cachedVersion); err != nil {
		return app.Version{}, err
	}
	var version app.Version
	err := c.getUnmarshal("/api/version", nil, &version)
	if err != nil {
		return app.Version{}, err
	}
	hadCheckedVersion := c.hasCheckedVersion
	c.hasCheckedVersion = true
	v, err := semver.ParseTolerant(version.Version)
	if err != nil {
		if !hadCheckedVersion {
			log.Warn().WithError(err).
				WithString("version", version.Version).
				Message("Failed to parse version from API. Version validation is turned off. Use at your own risk.")
		}
		return version, nil
	}
	c.cachedVersion = &v
	return version, nil
}
