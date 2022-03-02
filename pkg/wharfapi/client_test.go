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
			apiVer := semver.MustParse(test.apiVer)
			c := Client{ErrIfOutdatedServer: true, cachedVersion: &apiVer, hasCheckedVersion: true}
			endpointVer := semver.MustParse(test.endpointVer)
			err := c.validateEndpointVersion(endpointVer.Major, endpointVer.Minor, endpointVer.Patch)
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
			apiVer := semver.MustParse(test.apiVer)
			c := Client{ErrIfOutdatedServer: true, cachedVersion: &apiVer, hasCheckedVersion: true}
			endpointVer := semver.MustParse(test.endpointVer)
			err := c.validateEndpointVersion(endpointVer.Major, endpointVer.Minor, endpointVer.Patch)
			require.Error(t, err)
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
	oldHighest := HighestSupportedVersion
	defer func() { HighestSupportedVersion = oldHighest }()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			apiVer := semver.MustParse(test.apiVer)
			c := Client{ErrIfOutdatedClient: true, cachedVersion: &apiVer, hasCheckedVersion: true}
			HighestSupportedVersion = semver.MustParse(test.clientVer)
			err := c.validateEndpointVersion(apiVer.Major, apiVer.Minor, apiVer.Patch)
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
			apiVer := semver.MustParse(test.apiVer)
			c := Client{ErrIfOutdatedClient: true, cachedVersion: &apiVer, hasCheckedVersion: true}
			HighestSupportedVersion = semver.MustParse(test.clientVer)
			err := c.validateEndpointVersion(apiVer.Major, apiVer.Minor, apiVer.Patch)
			require.Error(t, err)
			require.ErrorIs(t, err, ErrOutdatedClient)
		})
	}
}
