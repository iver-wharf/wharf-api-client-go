// Package request contains plain old Go types used in the Gin endpoint handlers
// and Swaggo documentation for the HTTP request models, with Gin- and
// Swaggo-specific Go tags.
//
// Copied from https://github.com/iver-wharf/wharf-api/blob/v5.1.2/pkg/model/request/request.go
package request

import (
	"time"
)

// Reference doc about the Go tags:
//  TAG                  SOURCE                   DESCRIPTION
//  json:"foo"           encoding/json            Serializes field with the name "foo"
//  format:"date-time"   swaggo/swag              Swagger format
//  validate:"required"  swaggo/swag              Mark Swagger field as required/non-nullable
//  binding:"required"   go-playground/validator  Gin's Bind will error if nil or zero
//
// go-playground/validator uses the tag "validate" by default, but Gin overrides
// changes that to "binding".

// TokenSearch holds values used in verbatim searches for tokens.
type TokenSearch struct {
	Token    string `json:"token" format:"password"`
	UserName string `json:"userName"`
}

// Token specifies fields when creating a new token.
type Token struct {
	Token      string `json:"token" format:"password" validate:"required"`
	UserName   string `json:"userName" validate:"required"`
	ProviderID uint   `json:"providerId" minimum:"0"`
}

// TokenUpdate specifies fields when updating a token.
type TokenUpdate struct {
	Token    string `json:"token" format:"password" validate:"required"`
	UserName string `json:"userName" validate:"required"`
}

// Branch specifies fields when adding a new branch to a project.
type Branch struct {
	Name    string `json:"name" validate:"required"`
	Default bool   `json:"default"`
}

// BranchUpdate specifies fields for a single branch.
type BranchUpdate struct {
	Name string `json:"name" validate:"required"`
}

// BranchListUpdate specifies fields when resetting all branches for a project.
type BranchListUpdate struct {
	DefaultBranch string         `json:"defaultBranch" extensions:"x-nullable"`
	Branches      []BranchUpdate `json:"branches"`
}

// LogOrStatusUpdate is a single log line, together with its timestamp of when
// it was logged; or a build status update.
//
// The build status field takes precedence, and if set it will update the
// build's status, while the message and the timestamp is ignored.
type LogOrStatusUpdate struct {
	Message   string      `json:"message"`
	Timestamp time.Time   `json:"timestamp" format:"date-time"`
	Status    BuildStatus `json:"status" enums:",Scheduling,Running,Completed,Failed"`
}

// BuildStatus is an enum of different states for a build.
type BuildStatus string

const (
	// BuildScheduling means the build has been registered, but no code
	// execution has begun yet. This is usually quite an ephemeral state.
	BuildScheduling BuildStatus = "Scheduling"
	// BuildRunning means the build is executing right now. The execution
	// engine has load in the target code paths and repositories.
	BuildRunning BuildStatus = "Running"
	// BuildCompleted means the build has finished execution successfully.
	BuildCompleted BuildStatus = "Completed"
	// BuildFailed means that something went wrong with the build. Could be a
	// misconfiguration in the .wharf-ci.yml file, or perhaps a scripting error
	// in some build step.
	BuildFailed BuildStatus = "Failed"
)

// BuildStatusUpdate allows you to update the status of a build.
type BuildStatusUpdate struct {
	Status BuildStatus `json:"status" enums:"Scheduling,Running,Completed,Failed"`
}

// BuildInputs is a key-value object of input variables used when starting a new
// build, where the key is the input variable name and the value is its string,
// boolean, or numeric value.
type BuildInputs map[string]interface{}

// Project specifies fields when creating a new project.
type Project struct {
	Name            string `json:"name" validate:"required" binding:"required"`
	GroupName       string `json:"groupName"`
	Description     string `json:"description"`
	AvatarURL       string `json:"avatarUrl"`
	TokenID         uint   `json:"tokenId" minimum:"0"`
	ProviderID      uint   `json:"providerId" minimum:"0"`
	BuildDefinition string `json:"buildDefinition"`
	GitURL          string `json:"gitUrl"`
	RemoteProjectID string `json:"remoteProjectId"`
}

// ProjectUpdate specifies fields when updating a project.
type ProjectUpdate struct {
	Name            string `json:"name" validate:"required" binding:"required"`
	GroupName       string `json:"groupName"`
	Description     string `json:"description"`
	AvatarURL       string `json:"avatarUrl"`
	TokenID         uint   `json:"tokenId" minimum:"0"`
	ProviderID      uint   `json:"providerId" minimum:"0"`
	BuildDefinition string `json:"buildDefinition"`
	GitURL          string `json:"gitUrl"`
}

// ProjectOverridesUpdate specifies fields when updating a project's overrides.
type ProjectOverridesUpdate struct {
	Description string `json:"description"`
	AvatarURL   string `json:"avatarUrl"`
	GitURL      string `json:"gitUrl"`
}

// ProviderName is an enum of different providers that are available over at
// https://github.com/iver-wharf
type ProviderName string

const (
	// ProviderAzureDevOps refers to the Azure DevOps provider plugin,
	// https://github.com/iver-wharf/wharf-provider-azuredevops
	ProviderAzureDevOps ProviderName = "azuredevops"
	// ProviderGitLab refers to the GitLab provider plugin,
	// https://github.com/iver-wharf/wharf-provider-gitlab
	ProviderGitLab ProviderName = "gitlab"
	// ProviderGitHub refers to the GitHub provider plugin,
	// https://github.com/iver-wharf/wharf-provider-github
	ProviderGitHub ProviderName = "github"
	// ProviderNameValues is a concatenated list of the different provider names
	// available. Useful in validation error messages.
	ProviderNameValues = ProviderAzureDevOps + ", " + ProviderGitLab + ", " + ProviderGitHub
)

// IsValid returns false if the underlying type is an unknown enum value.
// 	ProviderGitHub.IsValid()     // => true
// 	(ProviderName("")).IsValid() // => false
func (name ProviderName) IsValid() bool {
	return name == ProviderAzureDevOps ||
		name == ProviderGitLab ||
		name == ProviderGitHub
}

// ValidString returns the name as a string if valid, as well as the boolean
// value true, or false if the name is invalid.
// 	ProviderGitHub.ValidString()     // => "github", true
// 	(ProviderName("")).ValidString() // => "", false
func (name ProviderName) ValidString() (string, bool) {
	if name.IsValid() {
		return string(name), true
	}
	return "", false
}

// ProviderSearch holds values used in verbatim searches for providers.
type ProviderSearch struct {
	Name    ProviderName `json:"name" enums:"azuredevops,gitlab,github"`
	URL     string       `json:"url"`
	TokenID uint         `json:"tokenId" minimum:"0"`
}

// Provider specifies fields when creating a new provider.
type Provider struct {
	Name    ProviderName `json:"name" enums:"azuredevops,gitlab,github" validate:"required" binding:"required"`
	URL     string       `json:"url" validate:"required" binding:"required"`
	TokenID uint         `json:"tokenId" minimum:"0"`
}

// ProviderUpdate specifies fields when updating a provider.
type ProviderUpdate struct {
	Name    ProviderName `json:"name" enums:"azuredevops,gitlab,github" validate:"required" binding:"required"`
	URL     string       `json:"url" validate:"required" binding:"required"`
	TokenID uint         `json:"tokenId" minimum:"0"`
}
