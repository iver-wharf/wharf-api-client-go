package wharfapi

import (
	"fmt"

	"github.com/iver-wharf/wharf-api-client-go/v2/pkg/model/response"
)

// GetBuildAllTestResultDetailList fetches all the test result
// details for the specified build by invoking the HTTP request:
//  GET /api/build/{buildId}/test-result/detail
//
// Added in wharf-api v5.0.0.
func (c *Client) GetBuildAllTestResultDetailList(buildID uint) (response.PaginatedTestResultDetails, error) {
	if err := c.validateEndpointVersion(5, 0, 0); err != nil {
		return response.PaginatedTestResultDetails{}, err
	}
	var details response.PaginatedTestResultDetails
	path := fmt.Sprintf("/api/build/%d/test-result/detail", buildID)
	err := c.getUnmarshal(path, nil, &details)
	return details, err
}

// GetBuildAllTestResultSummaryList fetches all the test result
// summaries for the specified build by invoking the HTTP request:
//  GET /api/build/{buildId}/test-result/summary
//
// Added in wharf-api v5.0.0.
func (c *Client) GetBuildAllTestResultSummaryList(buildID uint) (response.PaginatedTestResultSummaries, error) {
	if err := c.validateEndpointVersion(5, 0, 0); err != nil {
		return response.PaginatedTestResultSummaries{}, err
	}
	var summaries response.PaginatedTestResultSummaries
	path := fmt.Sprintf("/api/build/%d/test-result/summary", buildID)
	err := c.getUnmarshal(path, nil, &summaries)
	return summaries, err
}

// GetBuildTestResultSummary fetches a test result summary by ID by
// invoking the HTTP request:
//  GET /api/build/{buildId}/test-result/summary/{artifactId}
//
// Added in wharf-api v5.0.0.
func (c *Client) GetBuildTestResultSummary(buildID, artifactID uint) (response.TestResultSummary, error) {
	if err := c.validateEndpointVersion(5, 0, 0); err != nil {
		return response.TestResultSummary{}, err
	}
	var summary response.TestResultSummary
	path := fmt.Sprintf("/api/build/%d/test-result/summary/%d", buildID, artifactID)
	err := c.getUnmarshal(path, nil, &summary)
	return summary, err
}

// GetBuildTestResultDetailList fetches all test result details for the specified
// test result summary by invoking the HTTP request:
//  GET /api/build/{buildId}/test-result/summary/{artifactId}/detail
//
// Added in wharf-api v5.0.0.
func (c *Client) GetBuildTestResultDetailList(buildID, artifactID uint) (response.PaginatedTestResultDetails, error) {
	if err := c.validateEndpointVersion(5, 0, 0); err != nil {
		return response.PaginatedTestResultDetails{}, err
	}
	var details response.PaginatedTestResultDetails
	path := fmt.Sprintf("/api/build/%d/test-result/summary/%d/detail", buildID, artifactID)
	err := c.getUnmarshal(path, nil, &details)
	return details, err
}

// GetBuildAllTestResultListSummary fetches the test result list summary of all tests for
// the specified build.
//  GET /api/build/{buildId}/test-result/list-summary
//
// Added in wharf-api v5.0.0.
func (c *Client) GetBuildAllTestResultListSummary(buildID uint) (response.TestResultListSummary, error) {
	if err := c.validateEndpointVersion(5, 0, 0); err != nil {
		return response.TestResultListSummary{}, err
	}
	var listSummary response.TestResultListSummary
	path := fmt.Sprintf("/api/build%d/test-result/list-summary", buildID)
	err := c.getUnmarshal(path, nil, &listSummary)
	return listSummary, err
}
