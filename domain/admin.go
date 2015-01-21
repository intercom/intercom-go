package domain

import "fmt"

type Admin struct {
	Id    string
	Type  string
	Name  string
	Email string
}

func (a Admin) authorType() string {
	return "admin"
}

func (a Admin) IsNobodyAdmin() bool {
	return a.Type == "nobody_admin"
}

func (a Admin) String() string {
	return fmt.Sprintf("[intercom] %s { id: %s name: %s, email: %s }", a.Type, a.Id, a.Name, a.Email)
}
