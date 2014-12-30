package gojira

import (
	"bytes"
	"encoding/json"
	"net/url"
	"strconv"
	"time"
)

type Issue struct {
	Id        string
	Key       string
	Self      string
	Expand    string
	Fields    *IssueFields
	CreatedAt time.Time
}

type IssueList struct {
	Expand     string
	StartAt    int
	MaxResults int
	Total      int
	Issues     []*Issue
	Pagination *Pagination
}

type IssueFields struct {
	IssueType   *IssueType
	Summary     string
	Description string
	Reporter    *User
	Assignee    *User
	Project     *JiraProject
	Created     string
}

type IssueType struct {
	Self        string
	Id          string
	Description string
	IconUrl     string
	Name        string
	Subtask     bool
}

// Search for all issues assigned to a user.
func (j *Jira) IssuesAssignedTo(user string, maxResults int, startAt int) (i IssueList, err error) {
	uri := j.ApiPath + "/search?jql=assignee=\"" +
		url.QueryEscape(user) + "\"&startAt=" + strconv.Itoa(startAt) +
		"&maxResults=" + strconv.Itoa(maxResults)

	contents, err := j.getRequest(uri)
	if err != nil {
		return
	}

	var issues IssueList
	err = json.Unmarshal(contents, &issues)
	if err != nil {
		return
	}

	for _, issue := range issues.Issues {
		t, _ := time.Parse(dateLayout, issue.Fields.Created)
		issue.CreatedAt = t
	}

	pagination := Pagination{
		Total:      issues.Total,
		StartAt:    issues.StartAt,
		MaxResults: issues.MaxResults,
	}
	pagination.Compute()
	issues.Pagination = &pagination

	return issues, nil
}

// Returns an issue by it's id.
func (j *Jira) Issue(id string) (i Issue, err error) {
	uri := j.ApiPath + "/issue/" + id

	contents, err := j.getRequest(uri)
	if err != nil {
		return
	}

	var issue Issue
	err = json.Unmarshal(contents, &issue)
	if err != nil {
		return
	}

	return issue, nil
}

// Create a new comment on the issue.
func (j *Jira) AddComment(issue *Issue, comment string) error {
	var cMap = make(map[string]string)
	cMap["body"] = comment

	cJson, err := json.Marshal(cMap)
	if err != nil {
		return err
	}

	uri := j.BaseUrl + j.ApiPath + "/issue/" + issue.Key + "/comment"
	body := bytes.NewBuffer(cJson)

	_, err = j.postJson(uri, body)
	if err != nil {
		return err
	}

	return nil
}
