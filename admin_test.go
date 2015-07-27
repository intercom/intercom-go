package intercom

import "testing"

func TestNobodyAdmin(t *testing.T) {
	admin := Admin{Type: "nobody_admin", ID: "123"}
	if admin.IsNobodyAdmin() != true {
		t.Errorf("Nobody Admin not recognised")
	}
}

func TestAdminMessageAddress(t *testing.T) {
	admin := Admin{ID: "123", Name: "josler"}
	address := admin.MessageAddress()
	if address.ID != "123" {
		t.Errorf("Admin address did not have ID")
	}
	if address.Type != "admin" {
		t.Errorf("Admin address was not of type admin, was %s", address.Type)
	}
	if address.Email != "" && address.UserID != "" {
		t.Errorf("Admin address had Email/UserID")
	}
}

func TestAdminList(t *testing.T) {
	adminService := AdminService{Repository: TestAdminAPI{t: t}}
	adminList, _ := adminService.List()
	if adminList.Admins[0].ID != "213" {
		t.Errorf("Admin not found")
	}
}

type TestAdminAPI struct {
	t *testing.T
}

func (t TestAdminAPI) list() (AdminList, error) {
	return AdminList{Admins: []Admin{Admin{ID: "213"}}}, nil
}
