package jira_client2

import (
	"fmt"
	"net/http"
)


type GroupsType struct {
	Size            int        `json:"size,omitempty" structs:"size,omitempty"`
	Items           []string   `json:"items,omitempty" structs:"items,omitempty"`
}

type ApplicationRolesType struct {
	Size            int        `json:"size,omitempty" structs:"size,omitempty"`
	Items           []string   `json:"items,omitempty" structs:"items,omitempty"`
}


// User represents a JIRA user.
type User struct {
	Self            string     `json:"self,omitempty" structs:"self,omitempty"`
	Name            string     `json:"name,omitempty" structs:"name,omitempty"`
	Password        string     `json:"-"`
	Key             string     `json:"key,omitempty" structs:"key,omitempty"`
	EmailAddress    string     `json:"emailAddress,omitempty" structs:"emailAddress,omitempty"`
	AvatarUrls      AvatarUrls `json:"avatarUrls,omitempty" structs:"avatarUrls,omitempty"`
	DisplayName     string     `json:"displayName,omitempty" structs:"displayName,omitempty"`
	Active          bool       `json:"active,omitempty" structs:"active,omitempty"`
	Group        	GroupsType `json:"groups,omitempty" structs:"groups,omitempty"`
	TimeZone        string     `json:"timeZone,omitempty" structs:"timeZone,omitempty"`
	ApplicationRoles ApplicationRolesType   `json:"applicationRoles,omitempty" structs:"applicationRoles,omitempty"`
	ApplicationKeys []string   `json:"applicationKeys,omitempty" structs:"applicationKeys,omitempty"`
	Lokale          string     `json:"lokale,omitempty" structs:"lokale,omitempty"`
	Expand          string     `json:"expand,omitempty" structs:"lokale,omitempty"`
}

type CreateUser struct {
	Self            string     `json:"self,omitempty" structs:"self,omitempty"`
	Name            string     `json:"name,omitempty" structs:"name,omitempty"`
	Password        string     `json:"password,omitempty" structs:"password,omitempty"`
	EmailAddress    string     `json:"emailAddress,omitempty" structs:"emailAddress,omitempty"`
	DisplayName     string     `json:"displayName,omitempty" structs:"displayName,omitempty"`
	ApplicationKeys []string   `json:"applicationKeys,omitempty" structs:"applicationKeys,omitempty"`
}

// UserGroup represents the group list
type UserGroup struct {
	Self string `json:"self,omitempty" structs:"self,omitempty"`
	Name string `json:"name,omitempty" structs:"name,omitempty"`
}

type ActorsType struct {
	Id          int     `json:"id,omitempty" structs:"id,omitempty"`
	DisplayName string 	`json:"displayName,omitempty" structs:"displayName,omitempty"`
	Type 		string 	`json:"type,omitempty" structs:"type,omitempty"`
	Name 		string 	`json:"name,omitempty" structs:"name,omitempty"`
}

type ProjectRole struct {
	Self 		string 		`json:"self,omitempty" structs:"self,omitempty"`
	Name 		string 		`json:"name,omitempty" structs:"name,omitempty"`
	Id          int        	`json:"id,omitempty" structs:"id,omitempty"`
	Description string 		`json:"description,omitempty" structs:"description,omitempty"`
	Actors 		[]ActorsType   `json:"actors,omitempty" structs:"actors,omitempty"`
}

//func (s *UserService) Search(username string) (*[]User, *Response, error) {
type UserSearchResponseType struct {
	User       []string     `json:"user,omitempty" structs:"user,omitempty"`
}

func  (c *JIRAClient) UserSearch(username string) (*[]User, *http.Response) {
	u := fmt.Sprintf("/rest/api/2/user/search?username=%s", username)

	//payload := nil

	response := new([]User)
	_, res2  := c.doRequest("GET", u , nil, &response)

//	fmt.Println("res: " + string(res))

	return response, res2
}

func  (c *JIRAClient) UserGet(username string) (*User, *http.Response) {
	u := fmt.Sprintf("/rest/api/2/user?username=%s", username)

	response := new(User)
	_, res2  := c.doRequest("GET", u , nil, &response)

	//fmt.Println("res: " + string(res))

	return response, res2
}

func  (c *JIRAClient) UserCreate(user *CreateUser) (*User, *http.Response) {
	u := fmt.Sprintf("/rest/api/2/user")

	response := new(User)
	_, res2  := c.doRequest("POST", u , user, &response)

//	fmt.Println("res: " + string(res))

	return response, res2
}

/*
// Get gets user info from JIRA
//
// JIRA API docs: https://docs.atlassian.com/jira/REST/cloud/#api/2/user-getUser
func (s *UserService) Get(username string) (*User, *Response, error) {
	apiEndpoint := fmt.Sprintf("/rest/api/2/user?username=%s", username)
	req, err := s.client.NewRequest("GET", apiEndpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	user := new(User)
	//fmt.Println("apiEndpoint: " + apiEndpoint)
	resp, err := s.client.Do(req, user)
	if err != nil {
		return nil, resp, NewJiraError(resp, err)
	}
	return user, resp, nil
}

func (s *UserService) Search(username string) (*[]User, *Response, error) {
	apiEndpoint := fmt.Sprintf("/rest/api/2/user/search?username=%s", username)
	req, err := s.client.NewRequest("GET", apiEndpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	user := new([]User)
	//fmt.Println("apiEndpoint: " + apiEndpoint)
	resp, err := s.client.Do(req, user)
	if err != nil {
		return nil, resp, NewJiraError(resp, err)
	}
	return user, resp, nil
}


// Create creates an user in JIRA.
//
// JIRA API docs: https://docs.atlassian.com/jira/REST/cloud/#api/2/user-createUser
func (s *UserService) Create(user *User) (*User, *Response, error) {
	apiEndpoint := "/rest/api/2/user"
	req, err := s.client.NewRequest("POST", apiEndpoint, user)
	if err != nil {
		return nil, nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return nil, resp, err
	}

	responseUser := new(User)
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		e := fmt.Errorf("Could not read the returned data")
		return nil, resp, NewJiraError(resp, e)
	}
	err = json.Unmarshal(data, responseUser)
	if err != nil {
		e := fmt.Errorf("Could not unmarshall the data into struct")
		return nil, resp, NewJiraError(resp, e)
	}
	return responseUser, resp, nil
}
*/
