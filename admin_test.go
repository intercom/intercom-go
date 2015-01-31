package intercom

import "testing"

func TestNobodyAdmin(t *testing.T) {
	admin := Admin{Type: "nobody_admin", Id: "123"}
	if admin.IsNobodyAdmin() != true {
		t.Errorf("Nobody Admin not recognised")
	}
}
