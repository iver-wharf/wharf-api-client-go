package wharfapi

import (
	"testing"

	"github.com/blang/semver/v4"
	"github.com/stretchr/testify/require"
)

func TestValidateVersion_serverVersionOK(t *testing.T) {
	var tests = []struct {
		name        string
		apiVer      string
		endpointVer string
	}{
		{
			name:        "same",
			apiVer:      "5.0.0",
			endpointVer: "5.0.0",
		},
		{
			name:        "newer/major-bump",
			apiVer:      "6.0.0",
			endpointVer: "5.0.0",
		},
		{
			name:        "newer/minor-bump",
			apiVer:      "5.1.0",
			endpointVer: "5.0.0",
		},
		{
			name:        "newer/patch-bump",
			apiVer:      "5.0.1",
			endpointVer: "5.0.0",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := testValidateServerVersion(t, test.apiVer, test.endpointVer)
			require.NoError(t, err)
		})
	}
}

func TestValidateVersion_serverVersionError(t *testing.T) {
	var tests = []struct {
		name        string
		apiVer      string
		endpointVer string
	}{
		{
			name:        "major-bump",
			apiVer:      "5.0.0",
			endpointVer: "6.0.0",
		},
		{
			name:        "minor-bump",
			apiVer:      "5.0.0",
			endpointVer: "5.1.0",
		},
		{
			name:        "patch-bump",
			apiVer:      "5.0.0",
			endpointVer: "5.0.5",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := testValidateServerVersion(t, test.apiVer, test.endpointVer)
			require.ErrorIs(t, err, ErrOutdatedServer)
		})
	}
}

func TestValidateVersion_clientVersionOk(t *testing.T) {
	var tests = []struct {
		name      string
		apiVer    string
		clientVer string
	}{
		{
			name:      "single-major-bump",
			apiVer:    "6.0.0",
			clientVer: "5.0.0",
		},
		{
			name:      "minor-bump",
			apiVer:    "5.9.0",
			clientVer: "5.0.0",
		},
		{
			name:      "patch-bumps",
			apiVer:    "5.0.7",
			clientVer: "5.0.0",
		},
		{
			name:      "all-bump",
			apiVer:    "6.8.11",
			clientVer: "5.0.0",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := testValidateClientVersion(t, test.apiVer, test.clientVer)
			require.NoError(t, err)
		})
	}
}

func TestValidateVersion_clientVersionError(t *testing.T) {
	var tests = []struct {
		name      string
		apiVer    string
		clientVer string
	}{
		{
			name:      "major-bumps",
			apiVer:    "7.0.0",
			clientVer: "5.0.0",
		},
		{
			name:      "all-bump",
			apiVer:    "7.8.11",
			clientVer: "5.15.20",
		},
	}
	oldHighest := HighestSupportedVersion
	defer func() { HighestSupportedVersion = oldHighest }()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := testValidateClientVersion(t, test.apiVer, test.clientVer)
			require.ErrorIs(t, err, ErrOutdatedClient)
		})
	}
}

func testValidateServerVersion(t *testing.T, apiVerStr, endpointVerStr string) error {
	apiVer := testParseVersion(t, apiVerStr)
	endpointVer := testParseVersion(t, endpointVerStr)
	c := Client{ErrIfOutdatedServer: true, cachedVersion: &apiVer, hasCheckedVersion: true}
	return c.validateEndpointVersion(endpointVer.Major, endpointVer.Minor, endpointVer.Patch)
}

func testValidateClientVersion(t *testing.T, apiVerStr, clientVerStr string) error {
	oldHighest := HighestSupportedVersion
	defer func() { HighestSupportedVersion = oldHighest }()

	apiVer := testParseVersion(t, apiVerStr)
	clientVer := testParseVersion(t, clientVerStr)
	HighestSupportedVersion = clientVer
	c := Client{ErrIfOutdatedClient: true, cachedVersion: &apiVer, hasCheckedVersion: true}
	return c.validateEndpointVersion(apiVer.Major, apiVer.Minor, apiVer.Patch)
}

func testParseVersion(t *testing.T, str string) semver.Version {
	v, err := semver.Parse(str)
	require.NoErrorf(t, err, "parse version: %q", str)
	return v
}
