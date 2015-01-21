package domain

import "fmt"

type Note struct {
	Id    string
	User  User
	Admin Admin
	Body  string
}

func (n Note) String() string {
	return fmt.Sprintf("[intercom] note { id: %s, body: %s }", n.Id, n.Body)
}
