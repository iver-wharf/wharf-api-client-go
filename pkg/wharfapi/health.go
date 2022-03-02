package wharfapi

import "github.com/iver-wharf/wharf-api-client-go/v2/pkg/model/response"

// GetVersion gets the health of the API by invoking the
// HTTP request:
//  GET /api/health
func (c *Client) GetHealth() (response.HealthStatus, error) {
	var health response.HealthStatus
	err := c.getUnmarshal("/api/health", nil, &health)
	return health, err
}

// Ping pings, and hopefully you get a pong in return, by invoking the
// HTTP request:
//  GET /api/ping
func (c *Client) Ping() (response.Ping, error) {
	var ping response.Ping
	err := c.getUnmarshal("/api/ping", nil, &ping)
	return ping, err
}
