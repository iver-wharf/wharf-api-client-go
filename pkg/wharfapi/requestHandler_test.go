package wharfapi

import (
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
