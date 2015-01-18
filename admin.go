package intercom

import "fmt"

type Admin struct {
	*Resource
	AdminParams
	Type  string `json:"type"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type AdminParams struct {
	Id   interface{}
	Open *bool // used in finding conversations for an Admin
}

type adminIdentifiers struct {
	Id interface{} `json:"id,omitempty" url:"id,omitempty"`
}

func (a Admin) String() string {
	return fmt.Sprintf("[intercom] %s { id: %s name: %s, email: %s }", a.Type, a.Id, a.Name, a.Email)
}

type AdminList struct {
	Admins []Admin `json:"admins"`
}

func (a Admin) List() (*AdminList, error) {
	if responseBody, err := a.client.Get("/admins", nil); err != nil {
		return nil, err
	} else {
		adminList := AdminList{}
		return &adminList, a.Unmarshal(&adminList, responseBody.([]byte))
	}
}

func (a Admin) IsNobodyAdmin() bool {
	return a.Type == "nobody_admin"
}
