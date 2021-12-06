package wharfapi

import (
	"errors"
	"fmt"

	"github.com/iver-wharf/wharf-api/pkg/model/response"
)

// CreateBuildTestResult is not implemented yet.
// Should handle invoking the HTTP request:
//  POST /api/build/{buildId}/test-result
func (c Client) CreateBuildTestResult() error {
	return errors.New("not implemented yet")
}

// GetBuildAllTestResultDetailList fetches all the test result
// details for the specified build by invoking the HTTP request:
//  GET /api/build/{buildId}/test-result/detail
func (c Client) GetBuildAllTestResultDetailList(buildID uint) ([]response.TestResultDetail, error) {
	details := response.PaginatedTestResultDetails{}
	path := fmt.Sprintf("/api/build/%d/test-result/detail", buildID)
	err := c.GetDecoded(&details, "TEST RESULT", path, nil)
	return details.List, err
}

// GetBuildAllTestResultSummaryList fetches all the test result
// summaries for the specified build by invoking the HTTP request:
//  GET /api/build/{buildId}/test-result/summary
func (c Client) GetBuildAllTestResultSummaryList(buildID uint) ([]response.TestResultSummary, error) {
	summaries := response.PaginatedTestResultSummaries{}
	path := fmt.Sprintf("/api/build/%d/test-result/summary", buildID)
	err := c.GetDecoded(&summaries, "TEST RESULT", path, nil)
	return summaries.List, err
}

// GetBuildTestResultSummary fetches a test result summary by ID by
// invoking the HTTP request:
//  GET /api/build/{buildId}/test-result/summary/{artifactId}
func (c Client) GetBuildTestResultSummary(buildID, artifactID uint) (response.TestResultSummary, error) {
	summary := response.TestResultSummary{}
	path := fmt.Sprintf("/api/build/%d/test-result/summary/%d", buildID, artifactID)
	err := c.GetDecoded(&summary, "TEST RESULT", path, nil)
	return summary, err
}

// GetBuildTestResultDetailList fetches all test result details for the specified
// test result summary by invoking the HTTP request:
//  GET /api/build/{buildId}/test-result/summary/{artifactId}/detail
func (c Client) GetBuildTestResultDetailList(buildID, artifactID uint) ([]response.TestResultDetail, error) {
	details := response.PaginatedTestResultDetails{}
	path := fmt.Sprintf("/api/build/%d/test-result/summary/%d/detail", buildID, artifactID)
	err := c.GetDecoded(&details, "TEST RESULT", path, nil)
	return details.List, err
}

// GetBuildAllTestResultListSummary fetches the test result list summary of all tests for
// the specified build.
//  GET /api/build/{buildId}/test-result/list-summary
func (c Client) GetBuildAllTestResultListSummary(buildID uint) (response.TestResultListSummary, error) {
	listSummary := response.TestResultListSummary{}
	path := fmt.Sprintf("/api/build%d/test-result/list-summary", buildID)
	err := c.GetDecoded(&listSummary, "TEST RESULT", path, nil)
	return listSummary, err
}
