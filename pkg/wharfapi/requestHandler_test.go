package wharfapi

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveTokenFromJSON(t *testing.T) {
	type testCase struct {
		name string
		body string
		want string
	}

	tests := []testCase{
		{
			name: "Not matched",
			body: `{"tokenId":0,"userName":"","providerId":0}`,
			want: `{"tokenId":0,"userName":"","providerId":0}`,
		},
		{
			name: "Remove token from not formed JSON",
			body: `{"tokenId":0,"token":"some token","userName":"","providerId":0}`,
			want: `{"tokenId":0,"token":"*REDACTED*","userName":"","providerId":0}`,
		},
		{
			name: "Remove token from not formed JSON with white spaces",
			body: `{"tokenId" : 0,"token" : "some token","userName" : "","providerId" : 0}`,
			want: `{"tokenId" : 0,"token":"*REDACTED*","userName" : "","providerId" : 0}`,
		},
		{
			name: "Remove token from formed JSON",
			body: `{
    "tokenId": 0,
    "token": "some token",
    "userName": "",
    "providerId": 0
}
`,
			want: `{
    "tokenId": 0,
    "token":"*REDACTED*",
    "userName": "",
    "providerId": 0
}
`,
		},
		{
			name: "Remove token from invalid formed JSON",
			body: `{
    "tokenId": 0,
    "token":
"some token",
    "userName": "",
    "providerId": 0
}
`,
			want: `{
    "tokenId": 0,
    "token":"*REDACTED*",
    "userName": "",
    "providerId": 0
}
`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := redactTokenInJSON(tc.body)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestSanitizeURL(t *testing.T) {
	type testCase struct {
		name string
		url  string
		want string
	}

	tests := []testCase{
		{
			name: "Remove Token from URL",
			url:  "https:\\\\urlhere?Token=some123token&param=value&test=sth",
			want: "https:\\\\urlhere?Token=*REDACTED*&param=value&test=sth",
		},
		{
			name: "Remove token from URL",
			url:  "https:\\\\urlhere?test=sth&token=some123token",
			want: "https:\\\\urlhere?test=sth&token=*REDACTED*",
		},
		{
			name: "URL without token",
			url:  "https:\\\\urlhere?test=sth",
			want: "https:\\\\urlhere?test=sth",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := redactTokenInURL(tc.url)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestIsNonSuccessful_true(t *testing.T) {
	tests := []struct {
		name   string
		status int
	}{
		{
			name:   "100 Continue",
			status: http.StatusContinue,
		},
		{
			name:   "302 Found",
			status: http.StatusFound,
		},
		{
			name:   "404 Not Found",
			status: http.StatusNotFound,
		},
		{
			name:   "502 Bad Gateway",
			status: http.StatusBadGateway,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := isNonSuccessful(tc.status)
			assert.True(t, got)
		})
	}
}

func TestIsNonSuccessful_false(t *testing.T) {
	tests := []struct {
		name   string
		status int
	}{
		{
			name:   "200 OK",
			status: http.StatusOK,
		},
		{
			name:   "201 Created",
			status: http.StatusCreated,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := isNonSuccessful(tc.status)
			assert.False(t, got)
		})
	}
}
