package intercom

import (
	"encoding/json"
	"gopkg.in/intercom/intercom-go.v2/interfaces"
)

// AdminRepository defines the interface for working with Admins through the API.
type AdminRepository interface {
	list() (AdminList, error)
}

// AdminAPI implements AdminRepository
type AdminAPI struct {
	httpClient interfaces.HTTPClient
}

func (api AdminAPI) list() (AdminList, error) {
	adminList := AdminList{}
	data, err := api.httpClient.Get("/admins", nil)
	if err != nil {
		return adminList, err
	}
	err = json.Unmarshal(data, &adminList)
	return adminList, err
}
