package wharfapi

import "github.com/iver-wharf/wharf-api-client-go/v2/pkg/model/response"

// GetHealth gets the health of the API by invoking the
// HTTP request:
//  GET /api/health
//
// Added in wharf-api v0.7.1.
func (c *Client) GetHealth() (response.HealthStatus, error) {
	if err := c.validateEndpointVersion(0, 7, 1); err != nil {
		return response.HealthStatus{}, err
	}
	var health response.HealthStatus
	err := c.getUnmarshal("/api/health", nil, &health)
	return health, err
}

// Ping pings, and hopefully you get a pong in return, by invoking the
// HTTP request:
//  GET /api/ping
//
// Added in wharf-api v4.2.0.
func (c *Client) Ping() (response.Ping, error) {
	if err := c.validateEndpointVersion(4, 2, 0); err != nil {
		return response.Ping{}, err
	}
	var ping response.Ping
	err := c.getUnmarshal("/api/ping", nil, &ping)
	return ping, err
}
