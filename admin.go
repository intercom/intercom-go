package intercom

import (
	"encoding/json"
	"fmt"
)

type Admin struct {
	ID    json.Number `json:"id"`
	Type  string      `json:"type"`
	Name  string      `json:"name"`
	Email string      `json:"email"`
}

type AdminList struct {
	Admins []Admin
}

type AdminService struct {
	Repository AdminRepository
}

func (c *AdminService) List() (AdminList, error) {
	return c.Repository.list()
}

func (a Admin) IsNobodyAdmin() bool {
	return a.Type == "nobody_admin"
}

func (a Admin) String() string {
	return fmt.Sprintf("[intercom] %s { id: %s name: %s, email: %s }", a.Type, a.ID, a.Name, a.Email)
}
