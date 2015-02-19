package intercom

import "testing"

func TestNobodyAdmin(t *testing.T) {
	admin := Admin{Type: "nobody_admin", ID: "123"}
	if admin.IsNobodyAdmin() != true {
		t.Errorf("Nobody Admin not recognised")
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
