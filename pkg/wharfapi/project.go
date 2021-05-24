package wharfapi

import (
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type Project struct {
	ProjectID       uint   `json:"projectId"`
	Name            string `json:"name"`
	GroupName       string `json:"groupName"`
	BuildDefinition string `json:"buildDefinition"`
	TokenID         uint   `json:"tokenId"`
	Description     string `json:"description"`
	AvatarURL       string `json:"avatarUrl"`
	ProviderID      uint   `json:"providerId"`
	GitURL          string `json:"gitUrl"`
}

type ProjectRun struct {
	ProjectID   uint   `json:"projectId"`
	Stage       string `json:"stage"`
	Branch      string `json:"branch"`
	Environment string `json:"environment"`
}

type ProjectRunResponse struct {
	BuildID uint `json:"buildRef"`
}

func (c Client) GetProjectByID(projectID uint) (Project, error) {
	url := fmt.Sprintf("%s/api/project/%v", c.ApiUrl, projectID)
	ioBody, err := doRequest("GET | PROJECT |", http.MethodGet, url, []byte{}, c.AuthHeader)
	if err != nil {
		return Project{}, err
	}

	defer (*ioBody).Close()

	newProject := Project{}
	err = json.NewDecoder(*ioBody).Decode(&newProject)
	if err != nil {
		return Project{}, err
	}
	return newProject, nil
}

func (c Client) PutProject(project Project) (Project, error) {
	body, err := json.Marshal(project)
	if err != nil {
		return Project{}, err
	}

	log.WithField("project", string(body)).Traceln()

	url := fmt.Sprintf("%s/api/project", c.ApiUrl)
	ioBody, err := doRequest("PUT | PROJECT |", http.MethodPut, url, body, c.AuthHeader)
	if err != nil {
		return Project{}, err
	}

	defer (*ioBody).Close()

	newProject := Project{}
	err = json.NewDecoder(*ioBody).Decode(&newProject)
	if err != nil {
		return Project{}, err
	}

	return newProject, nil
}

func (c Client) PostProjectRun(projectRun ProjectRun) (ProjectRunResponse, error) {
	body, err := json.Marshal(projectRun)
	if err != nil {
		return ProjectRunResponse{}, err
	}

	url := fmt.Sprintf(
		"%s/api/project/%d/%s/run?branch=%s&environment=%s",
		c.ApiUrl,
		projectRun.ProjectID,
		projectRun.Stage,
		projectRun.Branch,
		projectRun.Environment)
	ioBody, err := doRequest("POST | PROJECT RUN |", http.MethodPut, url, body, c.AuthHeader)
	if err != nil {
		return ProjectRunResponse{}, err
	}

	defer (*ioBody).Close()

	newProject := ProjectRunResponse{}
	err = json.NewDecoder(*ioBody).Decode(&newProject)
	if err != nil {
		return ProjectRunResponse{}, err
	}

	log.WithField("ProjectRunResponse", newProject).Debugln()
	return newProject, nil
}
