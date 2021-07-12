package wharfapi

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
