package wharfapi

type AuthError struct {
	Realm string
}

func (e *AuthError) Error() string {
	return e.Realm
}

// WharfClient contains authentication and API URLs used to access
// the Wharf main API.
type Client struct {
	AuthHeader string
	ApiUrl     string
}

// WharfClient contains authentication and API URLs used to access
// the Wharf main API.
//
// Deprecated: This type has been renamed to Client and may be removed in a future release.
type WharfClient Client
