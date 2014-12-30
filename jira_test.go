package gojira

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNotFoundResource(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(http.NotFound))
	defer ts.Close()

	jira := NewJira(
		ts.URL,
		"/rest/api/2",
		"/activity",
		&Auth{"user", "password"},
	)

	url := fmt.Sprintf("%s/rest/api/2/user?username=john", ts.URL)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Error("error while creating a new request", err)
	}

	_, err = jira.execRequest(req)
	if err != ErrNotFound {
		t.Error("expected a not found error")
	}
}

func TestNotAuthorized(t *testing.T) {
	testHandler := func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "401 unauthorized", http.StatusUnauthorized)
	}

	ts := httptest.NewServer(http.HandlerFunc(testHandler))
	defer ts.Close()

	jira := NewJira(
		ts.URL,
		"/rest/api/2",
		"/activity",
		&Auth{"user", "password"},
	)

	url := fmt.Sprintf("%s/rest/api/2/user?username=john", ts.URL)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Error("error while creating a new request", err)
	}

	_, err = jira.execRequest(req)
	if err != ErrUnauthorized {
		t.Error("expected an unauthorized error")
	}
}
