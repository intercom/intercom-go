package intercom

import "fmt"

type Admin struct {
	*Resource
	Type  string `json:"type"`
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (a Admin) String() string {
	return fmt.Sprintf("[intercom] %s { id: %s name: %s, email: %s }", a.Type, a.Id, a.Name, a.Email)
}

type AdminList struct {
	Pages  Pages   `json:"pages,omitempty"`
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
