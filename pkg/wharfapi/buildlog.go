package wharfapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type BuildLog struct {
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

func (c Client) PostLog(buildID uint, buildLog BuildLog) error {
	body, err := json.Marshal(buildLog)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s/api/build/%d/log", c.ApiUrl, buildID)
	_, err = doRequest("POST | LOG", http.MethodPost, url, body, c.AuthHeader)
	if err != nil {
		return err
	}

	return nil
}

func (c Client) PutStatus(buildID uint, statusID BuildStatus) (Build, error) {
	uri := fmt.Sprintf("%s/api/build/%d?status=%s", c.ApiUrl, buildID, url.QueryEscape(statusID.String()))

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
