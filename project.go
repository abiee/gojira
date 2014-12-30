package gojira

type JiraProject struct {
	Self       string
	Id         string
	Key        string
	Name       string
	AvatarUrls map[string]string
}
