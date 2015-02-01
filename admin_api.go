package intercom

import (
	"encoding/json"

	"github.com/intercom/intercom-go/interfaces"
)

type AdminRepository interface {
	list() (AdminList, error)
}

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
