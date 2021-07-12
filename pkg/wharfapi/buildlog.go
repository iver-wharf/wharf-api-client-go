package wharfapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// BuildLog is a single log line from a given build. The Timestamp refers to
// when the logged line was printed inside the build.
type BuildLog struct {
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

// PostLog adds a new log to a build by invoking the HTTP request:
// 	POST /api/build/{buildID}/log
func (c Client) PostLog(buildID uint, buildLog BuildLog) error {
	body, err := json.Marshal(buildLog)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s/api/build/%d/log", c.APIURL, buildID)
	_, err = doRequest("POST | LOG", http.MethodPost, url, body, c.AuthHeader)
	if err != nil {
		return err
	}

	return nil
}

// PutStatus updates a build by invoking the HTTP request:
// 	PUT /api/build/{buildID}
func (c Client) PutStatus(buildID uint, statusID BuildStatus) (Build, error) {
	uri := fmt.Sprintf("%s/api/build/%d?status=%s", c.APIURL, buildID, url.QueryEscape(statusID.String()))

	ioBody, err := doRequest("PUT | STATUS", http.MethodPut, uri, nil, c.AuthHeader)
	if err != nil {
		return Build{}, err
	}
	defer (*ioBody).Close()

	build := Build{}
	err = json.NewDecoder(*ioBody).Decode(&build)
	if err != nil {
		return Build{}, err
	}

	return build, nil
}
