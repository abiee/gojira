package gojira

import (
    "fmt"
	"encoding/json"
)

const (
	PROJECT_URI = "/project"
)

type JiraProject struct {
	Self       string `json:"self"`
	Id         string `json:"id"`
	Key        string `json:"key"`
	Name       string `json:"name"`
	AvatarUrls map[string]string `json:"avatarUrls"`
}

// Returns a list of all available projects, aparently Atlassian does not
// allow to paginate the results.
func (j *Jira) GetAllProjects() (p []JiraProject, err error) {
    uri := fmt.Sprintf("%s%s", j.ApiPath, PROJECT_URI)
    contents, err := j.getRequest(uri)
    if err != nil {
        return
    }

    projects := make([]JiraProject, 0)
    err = json.Unmarshal(contents, &projects)
    if err != nil {
        return
    }

    return projects, nil
}

// Get a single project by it's id or project key.
func (j *Jira) GetProject(idOrKey string) (p *JiraProject, err error) {
    uri := fmt.Sprintf("%s%s/%s", j.ApiPath, PROJECT_URI, idOrKey)
    contents, err := j.getRequest(uri)
    if err != nil {
        return
    }

    project := new(JiraProject)
    err = json.Unmarshal(contents, &project)
    if err != nil {
        return
    }

    return project, nil
}
