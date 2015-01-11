package intercom

import "fmt"

// Author could represent a User or Admin
type Author struct {
	Type  string `json:"type"`
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (a Author) String() string {
	return fmt.Sprintf("[intercom] %s { id: %s name: %s, email: %s }", a.Type, a.Id, a.Name, a.Email)
}
