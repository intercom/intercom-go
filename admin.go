package intercom

import "fmt"

type Admin struct {
	ID    string `json:"id"`
	Type  string `json:"type"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (a Admin) IsNobodyAdmin() bool {
	return a.Type == "nobody_admin"
}

func (a Admin) String() string {
	return fmt.Sprintf("[intercom] %s { id: %s name: %s, email: %s }", a.Type, a.ID, a.Name, a.Email)
}
