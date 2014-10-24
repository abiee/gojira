package gojira

import (
	"code.google.com/p/gopass"
	"github.com/jbhat/go-jira-client"
)

func TestComment() {
	pass, err := gopass.GetPass("Pass: ")
	if err != nil {
		panic(err.Error())
	}
	jira := gojira.NewJira("https://jira.corp.ooyala.com", "/rest/api/latest", "/activity", &gojira.Auth{"jbhat", pass})
	issue, err := jira.Issue("APPSPLAT-840")
	err = jira.AddComment(&issue, "test jira plugin")
	if err != nil {
		panic(err.Error())
	}
}
