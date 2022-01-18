// Package response contains plain old Go types returned by wharf-web in the
// HTTP responses, with Swaggo-specific Go tags.
//
// Copied from github.com/iver-wharf/wharf-api/pkg/model/response
package response

import (
	"time"

	"gopkg.in/guregu/null.v4"
)

// TimeMetadata contains fields of when an object was created/added to the
// database, and when any field was last updated.
type TimeMetadata struct {
	UpdatedAt *time.Time `json:"updatedAt" format:"date-time" extensions:"x-nullable"`
	CreatedAt *time.Time `json:"createdAt" format:"date-time" extensions:"x-nullable"`
}

// ArtifactJSONFields holds the JSON field names for each field.
// Useful in ordering statements to map the correct field to the correct
// database column.
var ArtifactJSONFields = struct {
	ArtifactID string
	Name       string
	FileName   string
}{
	ArtifactID: "artifactId",
	Name:       "name",
	FileName:   "fileName",
}

// Artifact holds the binary data as well as metadata about that binary such as
// the file name and which build it belongs to.
type Artifact struct {
	TimeMetadata
	ArtifactID uint   `json:"artifactId" minimum:"0"`
	BuildID    uint   `json:"buildId" minimum:"0"`
	Name       string `json:"name"`
	FileName   string `json:"fileName"`
}

// ArtifactMetadata contains the file name and artifact ID of an Artifact.
type ArtifactMetadata struct {
	TimeMetadata
	FileName   string `json:"fileName"`
	ArtifactID uint   `json:"artifactId" minimum:"0"`
}

// Branch holds details about a project's branch.
type Branch struct {
	TimeMetadata
	BranchID  uint   `json:"branchId" minimum:"0"`
	ProjectID uint   `json:"projectId" minimum:"0"`
	Name      string `json:"name"`
	Default   bool   `json:"default"`
	TokenID   uint   `json:"tokenId" minimum:"0"`
}

// BranchList holds a list of branches, and a separate field for the default
// branch (if any).
type BranchList struct {
	DefaultBranch *Branch  `json:"defaultBranch" extensions:"x-nullable"`
	Branches      []Branch `json:"branches"`
}

// BuildJSONFields holds the JSON field names for each field.
// Useful in ordering statements to map the correct field to the correct
// database column.
var BuildJSONFields = struct {
	BuildID     string
	Environment string
	CompletedOn string
	ScheduledOn string
	StartedOn   string
	Stage       string
	StatusID    string
	IsInvalid   string
}{
	BuildID:     "buildId",
	Environment: "environment",
	CompletedOn: "finishedOn",
	ScheduledOn: "scheduledOn",
	StartedOn:   "startedOn",
	Stage:       "stage",
	StatusID:    "statusId",
	IsInvalid:   "isInvalid",
}

// Build holds data about the state of a build. Which parameters was used to
// start it, what status it holds, et.al.
type Build struct {
	TimeMetadata
	BuildID               uint                  `json:"buildId" minimum:"0"`
	StatusID              int                   `json:"statusId" enums:"0,1,2,3"`
	Status                BuildStatus           `json:"status" enums:"Scheduling,Running,Completed,Failed"`
	ProjectID             uint                  `json:"projectId" minimum:"0"`
	ScheduledOn           null.Time             `json:"scheduledOn" format:"date-time" extensions:"x-nullable"`
	StartedOn             null.Time             `json:"startedOn" format:"date-time" extensions:"x-nullable"`
	CompletedOn           null.Time             `json:"finishedOn" format:"date-time" extensions:"x-nullable"`
	GitBranch             string                `json:"gitBranch"`
	Environment           null.String           `json:"environment" swaggertype:"string" extensions:"x-nullable"`
	Stage                 string                `json:"stage"`
	Params                []BuildParam          `json:"params"`
	IsInvalid             bool                  `json:"isInvalid"`
	TestResultSummaries   []TestResultSummary   `json:"testResultSummaries"`
	TestResultListSummary TestResultListSummary `json:"testResultListSummary"`
}

// BuildParam holds the name and value of an input parameter fed into a build.
type BuildParam struct {
	BuildID uint   `json:"buildId" minimum:"0"`
	Name    string `json:"name"`
	Value   string `json:"value"`
}

// BuildReferenceWrapper holds a build reference. A unique identifier to a
// build.
type BuildReferenceWrapper struct {
	BuildReference string `json:"buildRef" example:"123"`
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

// HealthStatus holds a human-readable string stating the health of the API and
// its integrations, as well as a boolean for easy machine-readability.
type HealthStatus struct {
	Message   string `json:"message" example:"API is healthy."`
	IsHealthy bool   `json:"isHealthy" example:"true"`
}

// Log is a single logged line for a build.
type Log struct {
	LogID     uint      `json:"logId" minimum:"0"`
	BuildID   uint      `json:"buildId" minimum:"0"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp" format:"date-time"`
}

// PaginatedArtifacts is a list of artifacts as well as the explicit total count
// field.
type PaginatedArtifacts struct {
	List       []Artifact `json:"list"`
	TotalCount int64      `json:"totalCount"`
}

// PaginatedBranches is a list of branches as well as an explicit total count
// field.
type PaginatedBranches struct {
	List          []Branch `json:"list"`
	TotalCount    int64    `json:"totalCount"`
	DefaultBranch *Branch  `json:"defaultBranch"`
}

// PaginatedBuilds is a list of builds as well as an explicit total count field.
type PaginatedBuilds struct {
	List       []Build `json:"list"`
	TotalCount int64   `json:"totalCount"`
}

// PaginatedProjects is a list of projects as well as the explicit total count
// field.
type PaginatedProjects struct {
	List       []Project `json:"list"`
	TotalCount int64     `json:"totalCount"`
}

// PaginatedTokens is a list of tokens as well as the explicit total count
// field.
type PaginatedTokens struct {
	List       []Token `json:"list"`
	TotalCount int64   `json:"totalCount"`
}

// PaginatedProviders is a list of providers as well as the explicit total count
// field.
type PaginatedProviders struct {
	List       []Provider `json:"list"`
	TotalCount int64      `json:"totalCount"`
}

// PaginatedTestResultDetails is a list of test result details as well as the
// explicit total count field.
type PaginatedTestResultDetails struct {
	List       []TestResultDetail `json:"list"`
	TotalCount int64              `json:"totalCount"`
}

// PaginatedTestResultSummaries is a list of test result summaries as well as
// the explicit total count field.
type PaginatedTestResultSummaries struct {
	List       []TestResultSummary `json:"list"`
	TotalCount int64               `json:"totalCount"`
}

// Ping pongs.
type Ping struct {
	Message string `json:"message" example:"pong"`
}

// ProjectJSONFields holds the JSON field names for each field.
// Useful in ordering statements to map the correct field to the correct
// database column.
var ProjectJSONFields = struct {
	ProjectID       string
	RemoteProjectID string
	Name            string
	GroupName       string
	Description     string
	GitURL          string
}{
	ProjectID:       "projectId",
	RemoteProjectID: "remoteProjectId",
	Name:            "name",
	GroupName:       "groupName",
	Description:     "description",
	GitURL:          "gitUrl",
}

// Project holds details about a project.
type Project struct {
	TimeMetadata
	ProjectID             uint        `json:"projectId" minimum:"0"`
	RemoteProjectID       string      `json:"remoteProjectId"`
	Name                  string      `json:"name"`
	GroupName             string      `json:"groupName"`
	Description           string      `json:"description"`
	AvatarURL             string      `json:"avatarUrl"`
	TokenID               uint        `json:"tokenId" minimum:"0"`
	ProviderID            uint        `json:"providerId" minimum:"0"`
	Provider              *Provider   `json:"provider" extensions:"x-nullable"`
	BuildDefinition       string      `json:"buildDefinition"`
	Branches              []Branch    `json:"branches"`
	GitURL                string      `json:"gitUrl"`
	ParsedBuildDefinition interface{} `json:"build" swaggertype:"object" extensions:"x-nullable"`
}

// ProjectOverrides holds field overrides for a project.
type ProjectOverrides struct {
	ProjectID   uint   `json:"projectId" minimum:"0"`
	Description string `json:"description"`
	AvatarURL   string `json:"avatarUrl"`
	GitURL      string `json:"gitUrl"`
}

// ProviderJSONFields holds the JSON field names for each field.
// Useful in ordering statements to map the correct field to the correct
// database column.
var ProviderJSONFields = struct {
	ProviderID string
	Name       string
	URL        string
	TokenID    string
}{
	ProviderID: "providerId",
	Name:       "name",
	URL:        "url",
	TokenID:    "tokenId",
}

// Provider holds metadata about a connection to a remote provider. Some of
// importance are the URL field of where to find the remote, and the token field
// used to authenticate.
type Provider struct {
	TimeMetadata
	ProviderID uint         `json:"providerId" minimum:"0"`
	Name       ProviderName `json:"name" enums:"azuredevops,gitlab,github"`
	URL        string       `json:"url"`
	TokenID    uint         `json:"tokenId" minimum:"0"`
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
)

// TestStatus is an enum of different states a test run or test summary can be
// in.
type TestStatus string

const (
	// TestStatusSuccess means the test run or test summary passed, or in the
	// case that there are multiple tests then that there are no failing tests
	// and at least one successful test.
	TestStatusSuccess TestStatus = "Success"

	// TestStatusFailed means the test run or test summary failed, or in the
	// case that there are multiple tests then that at least one test failed.
	TestStatusFailed TestStatus = "Failed"

	// TestStatusNoTests means the test run or test summary is inconclusive,
	// where there are neither any passing nor failing tests.
	TestStatusNoTests TestStatus = "No tests"
)

// TestResultDetail contains data about a single test in a test result file.
type TestResultDetail struct {
	TimeMetadata
	TestResultDetailID uint             `json:"testResultDetailId" minimum:"0"`
	ArtifactID         uint             `json:"artifactId" minimum:"0"`
	BuildID            uint             `json:"buildId" minimum:"0"`
	Name               string           `json:"name"`
	Message            null.String      `json:"message" swaggertype:"string" extensions:"x-nullable"`
	StartedOn          null.Time        `json:"startedOn" format:"date-time" extensions:"x-nullable"`
	CompletedOn        null.Time        `json:"completedOn" format:"date-time" extensions:"x-nullable"`
	Status             TestResultStatus `json:"status" enums:"Failed,Passed,Skipped"`
}

// TestResultListSummary contains data about several test result files.
type TestResultListSummary struct {
	BuildID uint `json:"buildId" minimum:"0"`
	Total   uint `json:"total"`
	Failed  uint `json:"failed"`
	Passed  uint `json:"passed"`
	Skipped uint `json:"skipped"`
}

// TestResultStatus is an enum of different states a test result can be in.
type TestResultStatus string

const (
	// TestResultStatusSuccess means the test succeeded.
	TestResultStatusSuccess TestResultStatus = "Success"
	// TestResultStatusFailed means the test failed.
	TestResultStatusFailed TestResultStatus = "Failed"
	// TestResultStatusSkipped means the test was skipped.
	TestResultStatusSkipped TestResultStatus = "Skipped"
)

// TestResultSummary contains data about a single test result file.
type TestResultSummary struct {
	TimeMetadata
	TestResultSummaryID uint   `json:"testResultSummaryId" minimum:"0"`
	FileName            string `json:"fileName"`
	ArtifactID          uint   `json:"artifactId" minimum:"0"`
	BuildID             uint   `json:"buildId" minimum:"0"`
	Total               uint   `json:"total"`
	Failed              uint   `json:"failed"`
	Passed              uint   `json:"passed"`
	Skipped             uint   `json:"skipped"`
}

// TestsResults holds how many builds has passed and failed. A test result has
// the status of "Failed" if there are any failed tests, "Success" if there are
// any passing tests and no failed tests, and "No tests" if there are no failed
// nor passing tests.
type TestsResults struct {
	Passed uint       `json:"passed"`
	Failed uint       `json:"failed"`
	Status TestStatus `json:"status" enums:"Success,Failed,No tests"`
}

// TokenJSONFields holds the JSON field names for each field.
// Useful in ordering statements to map the correct field to the correct
// database column.
var TokenJSONFields = struct {
	TokenID  string
	Token    string
	UserName string
}{
	TokenID:  "tokenId",
	Token:    "token",
	UserName: "userName",
}

// Token holds credentials for a remote provider.
type Token struct {
	TimeMetadata
	TokenID  uint   `json:"tokenId" minimum:"0"`
	Token    string `json:"token" format:"password"`
	UserName string `json:"userName"`
}
