package gojira

import (
	"fmt"
	"encoding/json"
)

const (
	USER_URI        = "/user"
	USER_SEARCH_URI = "/user/search"
	// http://example.com:8080/jira/rest/api/2/user/assignable/multiProjectSearch [GET]
	// http://example.com:8080/jira/rest/api/2/user/assignable/search [GET]
	// http://example.com:8080/jira/rest/api/2/user/avatar [POST, PUT]
	// http://example.com:8080/jira/rest/api/2/user/avatar/temporary [POST, POST]
	// http://example.com:8080/jira/rest/api/2/user/avatar/{id} [DELETE]
	// http://example.com:8080/jira/rest/api/2/user/avatars [GET]
	// http://example.com:8080/jira/rest/api/2/user/picker [GET]
	// http://example.com:8080/jira/rest/api/2/user/viewissue/search [GET]
)

type User struct {
	Self         string            `json:"self"`
	Name         string            `json:"name"`
	EmailAddress string            `json:"emailAddress"`
	DisplayName  string            `json:"displayName"`
	Active       bool              `json:"active"`
	TimeZone     string            `json:"timeZone"`
	AvatarUrls   map[string]string `json:"avatarUrls"`
	Expand       string            `json:"expand"`
	// "groups": {
	//     "size": 3,
	//     "items": [
	//         {
	//             "name": "jira-user",
	//             "self": "http://www.example.com/jira/rest/api/2/group?groupname=jira-user"
	//         },
	//         {
	//             "name": "jira-admin",
	//             "self": "http://www.example.com/jira/rest/api/2/group?groupname=jira-admin"
	//         },
	//         {
	//             "name": "important",
	//             "self": "http://www.example.com/jira/rest/api/2/group?groupname=important"
	//         }
	//     ]
	// }
}

// Returns a user by it's username. Hits the endpoint.
//
//  GET http://example.com:8080/jira/rest/api/2/user?username=USERNAME
func (j *Jira) User(username string) (u *User, err error) {
	uri := fmt.Sprintf("%s%s?username=%s", j.ApiPath, USER_URI, username)
	contents, err := j.getRequest(uri)
	if err != nil {
		return
	}

	user := new(User)
	err = json.Unmarshal(contents, &user)
	if err != nil {
		return
	}

	return user, nil
}

// Returns a list of users that match the search string. Hits the endpoint.
//
// 	GET http://example.com:8080/jira/rest/api/2/user/search
func (j *Jira) SearchUser(username string, startAt int, maxResults int, includeActive bool, includeInactive bool) (c []byte, err error) {
	uri := j.BaseUrl + j.ApiPath + USER_URI + "?username=" + username
	contents, err := j.getRequest(uri)
	if err != nil {
		return nil, err
	}
	// @todo
	return contents, nil
}
