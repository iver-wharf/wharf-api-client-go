package wharfapi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/iver-wharf/wharf-api-client-go/pkg/wharfapi/query"
	"github.com/iver-wharf/wharf-api/pkg/model/response"
)

type Project response.Project
type ProjectSearch struct {
	OrderBy          []string `query:"name:orderBy"`
	Limit            *int     `query:"name:limit"`
	Offset           *int     `query:"name:offset,requires:Limit,min:0"`
	Name             *string  `query:"name:name"`
	GroupName        *string  `query:"name:groupName"`
	Description      *string  `query:"name:description"`
	TokenID          *uint    `query:"name:tokenId"`
	ProviderID       *uint    `query:"name:providerId"`
	GitURL           *string  `query:"name:gitUrl"`
	NameMatch        *string  `query:"name:nameMatch,excluded_with:Name"`
	GroupNameMatch   *string  `query:"name:groupNameMatch,excluded_with:GroupName"`
	DescriptionMatch *string  `query:"name:descriptionMatch,excluded_with:Description"`
	GitURLMatch      *string  `query:"name:gitUrlMatch,excluded_with:GitURL"`
	Match            *string  `query:"name:match"`
}

// ProjectRun is a range of options you start a build with. The ProjectID and
// Stage fields are required when starting a build.
type ProjectRun struct {
	ProjectID   uint   `json:"projectId"`
	Stage       string `json:"stage"`
	Branch      string `json:"branch" query:"name:branch"`
	Environment string `json:"environment" query:"name:environment"`
}

// ProjectRunResponse contains metadata about the newly started build.
type ProjectRunResponse struct {
	BuildID uint `json:"buildRef"`
}

// GetProjectByID fetches a project by ID by invoking the HTTP request:
// 	GET /api/project/{projectID}
func (c Client) GetProjectByID(projectID uint) (Project, error) {
	path := fmt.Sprintf("/api/project/%v", projectID)
	ioBody, err := doRequestNew("GET | PROJECT |", http.MethodGet, c.APIURL, path, nil, []byte{}, c.AuthHeader)
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

// SearchProject tries to match the given project on all non-zero fields by
// invoking the HTTP request:
// 	POST /api/projects/search
func (c Client) GetProjectList(params ProjectSearch) ([]Project, error) {
	q, err := query.FromObj(params)
	if err != nil {
		return nil, fmt.Errorf("failed constructing query from object: %w", err)
	}

	ioBody, err := doRequestNew("GET | PROJECT LIST |", http.MethodGet, c.APIURL, "/api/project", q, nil, c.AuthHeader)
	if err != nil {
		return nil, err
	}

	defer (*ioBody).Close()

	var foundProjects []Project
	err = json.NewDecoder(*ioBody).Decode(&foundProjects)
	if err != nil {
		return nil, err
	}

	return foundProjects, nil
}

// PutProject tries to match an existing project by ID or name+group and updates
// it, or adds a new a project if none matched, by invoking the HTTP request:
// 	PUT /api/project
func (c Client) PutProject(project Project) (Project, error) {
	body, err := json.Marshal(project)
	if err != nil {
		return Project{}, err
	}

	ioBody, err := doRequestNew("PUT | PROJECT |", http.MethodPut, c.APIURL, "/api/project", nil, body, c.AuthHeader)
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

// PostProjectRun starts a new build by invoking the HTTP request:
// 	POST /api/project/{projectID}/{stage}/run
func (c Client) PostProjectRun(projectRun ProjectRun) (ProjectRunResponse, error) {
	body, err := json.Marshal(projectRun)
	if err != nil {
		return ProjectRunResponse{}, err
	}

	q, err := query.FromObj(projectRun)
	if err != nil {
		return ProjectRunResponse{}, err
	}

	path := fmt.Sprintf("/api/project/%d/%s/run", projectRun.ProjectID, projectRun.Stage)
	ioBody, err := doRequestNew("POST | PROJECT RUN |", http.MethodPut, c.APIURL, path, q, body, c.AuthHeader)
	if err != nil {
		return ProjectRunResponse{}, err
	}

	defer (*ioBody).Close()

	newProject := ProjectRunResponse{}
	err = json.NewDecoder(*ioBody).Decode(&newProject)
	if err != nil {
		return ProjectRunResponse{}, err
	}

	log.Debug().WithUint("buildRef", newProject.BuildID).Message("Started build.")
	return newProject, nil
}
