package wharfapi

import "github.com/iver-wharf/wharf-core/pkg/app"

// GetVersion gets the version of the API by invoking the
// HTTP request:
//  GET /api/version
func (c Client) GetVersion() (app.Version, error) {
	var version app.Version
	err := c.getUnmarshal("/api/version", nil, &version)
	return version, err
}
