package intercom

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/launchpadcentral/intercom-go/interfaces"
)

// AdminRepository defines the interface for working with Admins through the API.
type AdminRepository interface {
	list() (AdminList, error)
	find(AdminIdentifiers) (Admin, error)
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

func (api AdminAPI) find(params AdminIdentifiers) (Admin, error) {
	return unmarshalToAdmin(api.getClientForFind(params))
}

func unmarshalToAdmin(data []byte, err error) (Admin, error) {
	savedAdmin := Admin{}
	if err != nil {
		return savedAdmin, err
	}
	err = json.Unmarshal(data, &savedAdmin)
	return savedAdmin, err
}

func (api AdminAPI) getClientForFind(params AdminIdentifiers) ([]byte, error) {
	switch {
	case params.ID != "":
		return api.httpClient.Get(fmt.Sprintf("/admins/%s", params.ID), nil)
	case params.Email != "":
		return api.httpClient.Get("/admins", params)
	}
	return nil, errors.New("Missing Admin Identifier")
}
