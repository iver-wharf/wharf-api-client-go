package wharfapi

import (
	"testing"

	"github.com/Masterminds/semver/v3"
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
			apiVer:      "v5.0.0",
			endpointVer: "v5.0.0",
		},
		{
			name:        "newer/major-bump",
			apiVer:      "v6.0.0",
			endpointVer: "v5.0.0",
		},
		{
			name:        "newer/minor-bump",
			apiVer:      "v5.1.0",
			endpointVer: "v5.0.0",
		},
		{
			name:        "newer/patch-bump",
			apiVer:      "v5.0.1",
			endpointVer: "v5.0.0",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			apiVer := semver.MustParse(test.apiVer)
			c := Client{ErrIfOutdatedServer: true, cachedVersion: apiVer, hasCheckedVersion: true}
			endpointVer := semver.MustParse(test.endpointVer)
			err := c.validateEndpointVersion(endpointVer)
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
			apiVer:      "v5.0.0",
			endpointVer: "v6.0.0",
		},
		{
			name:        "minor-bump",
			apiVer:      "v5.0.0",
			endpointVer: "v5.1.0",
		},
		{
			name:        "patch-bump",
			apiVer:      "v5.0.0",
			endpointVer: "v5.0.5",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			apiVer := semver.MustParse(test.apiVer)
			c := Client{ErrIfOutdatedServer: true, cachedVersion: apiVer, hasCheckedVersion: true}
			endpointVer := semver.MustParse(test.endpointVer)
			err := c.validateEndpointVersion(endpointVer)
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
			apiVer:    "v6.0.0",
			clientVer: "v5.0.0",
		},
		{
			name:      "minor-bump",
			apiVer:    "v5.9.0",
			clientVer: "v5.0.0",
		},
		{
			name:      "patch-bumps",
			apiVer:    "v5.0.7",
			clientVer: "v5.0.0",
		},
		{
			name:      "all-bump",
			apiVer:    "v6.8.11",
			clientVer: "v5.0.0",
		},
	}
	oldHighest := HighestSupportedVersion
	defer func() { HighestSupportedVersion = oldHighest }()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			apiVer := semver.MustParse(test.apiVer)
			c := Client{ErrIfOutdatedClient: true, cachedVersion: apiVer, hasCheckedVersion: true}
			HighestSupportedVersion = semver.MustParse(test.clientVer)
			err := c.validateEndpointVersion(apiVer)
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
			apiVer:    "v7.0.0",
			clientVer: "v5.0.0",
		},
		{
			name:      "all-bump",
			apiVer:    "v7.8.11",
			clientVer: "v5.15.20",
		},
	}
	oldHighest := HighestSupportedVersion
	defer func() { HighestSupportedVersion = oldHighest }()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			apiVer := semver.MustParse(test.apiVer)
			c := Client{ErrIfOutdatedClient: true, cachedVersion: apiVer, hasCheckedVersion: true}
			HighestSupportedVersion = semver.MustParse(test.clientVer)
			err := c.validateEndpointVersion(apiVer)
			require.Error(t, err)
			require.ErrorIs(t, err, ErrOutdatedClient)
		})
	}
}
