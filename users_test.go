package gojira

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAnUser(t *testing.T) {
	testHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"name": "dennis"}`)
	}

	ts := httptest.NewServer(http.HandlerFunc(testHandler))
	defer ts.Close()

	jira := NewJira(
		ts.URL,
		"/rest/api/2",
		"/activity",
		&Auth{"user", "password"},
	)

	user, err := jira.User("john")

	if user.Name != "dennis" {
		t.Errorf("got %s; want %s", user.Name, "dennis")
	}

	if err != nil {
		t.Error(err)
	}
}
