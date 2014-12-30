package gojira

import (
	"bytes"
	"errors"
	"io/ioutil"
	"math"
	"net/http"
)

type Jira struct {
	BaseUrl      string
	ApiPath      string
	ActivityPath string
	Client       *http.Client
	Auth         *Auth
}

type Auth struct {
	Login    string
	Password string
}

type Pagination struct {
	Total      int
	StartAt    int
	MaxResults int
	Page       int
	PageCount  int
	Pages      []int
}

var ErrNotFound = errors.New("gojira: resource not found")
var ErrUnauthorized = errors.New("gojira: not authorized")

func (p *Pagination) Compute() {
	p.PageCount = int(math.Ceil(float64(p.Total) / float64(p.MaxResults)))
	p.Page = int(math.Ceil(float64(p.StartAt) / float64(p.MaxResults)))

	p.Pages = make([]int, p.PageCount)
	for i := range p.Pages {
		p.Pages[i] = i
	}
}

const (
	dateLayout = "2006-01-02T15:04:05.000-0700"
)

func NewJira(baseUrl string, apiPath string, activityPath string, auth *Auth) *Jira {
	client := &http.Client{}

	return &Jira{
		BaseUrl:      baseUrl,
		ApiPath:      baseUrl + apiPath,
		ActivityPath: activityPath,
		Client:       client,
		Auth:         auth,
	}
}

func (j *Jira) getRequest(uri string) ([]byte, error) {
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}
	return j.execRequest(req)
}

func (j *Jira) postJson(uri string, body *bytes.Buffer) ([]byte, error) {
	req, err := http.NewRequest("POST", uri, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return j.execRequest(req)
}

func (j *Jira) execRequest(req *http.Request) ([]byte, error) {
	req.SetBasicAuth(j.Auth.Login, j.Auth.Password)

	resp, err := j.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		return nil, ErrNotFound
	}

	if resp.StatusCode == 401 {
		return nil, ErrUnauthorized
	}

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return contents, nil
}
