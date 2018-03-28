package intercom

import (
	"encoding/json"
	"errors"
	"fmt"

	"gopkg.in/intercom/intercom-go.v2/interfaces"
)

// VisitorRepository defines the interface for working with Visitors through the API.
type VisitorRepository interface {
	find(UserIdentifiers) (Visitor, error)
	update(*Visitor) (Visitor, error)
	convert(*Visitor, *Lead) (Lead, error)
	convert(*Visitor, *User) (User, error)
	delete(id string) (Visitor, error)
}

// VisitorAPI implements VisitorRepository
type VisitorAPI struct {
	httpClient interfaces.HTTPClient
}

func (api VisitorAPI) find(params UserIdentifiers) (Visitor, error) {
	return unmarshalToVisitor(api.getClientForFind(params))
}

func (api VisitorAPI) getClientForFind(params UserIdentifiers) ([]byte, error) {
	switch {
	case params.ID != "":
		return api.httpClient.Get(fmt.Sprintf("/visitors/%s", params.ID), nil)
	case params.UserID != "":
		return api.httpClient.Get("/visitors", params)
	}
	return nil, errors.New("Missing Visitor Identifier")
}

func (api VisitorAPI) update(visitor *Visitor) (Visitor, error) {
	requestVisitor := api.buildRequestVisitor(visitor)
	return unmarshalToVisitor(api.httpClient.Post("/visitors", &requestVisitor))
}

func (api VisitorAPI) convert(visitor *Visitor, lead *Lead) (Lead, error) {
	cr := convertToLeadRequest{Visitor: api.buildRequestVisitor(visitor), Lead: requestUser{
		ID:         lead.ID,
		UserID:     lead.UserID,
		Email:      lead.Email,
		SignedUpAt: lead.SignedUpAt,
	}}
	return unmarshalToUser(api.httpClient.Post("/visitors/convert", &cr))
}

func (api VisitorAPI) convert(visitor *Visitor, user *User) (User, error) {
	cr := convertToUserRequest{Visitor: api.buildRequestVisitor(visitor), User: requestUser{
		ID:         user.ID,
		UserID:     user.UserID,
		Email:      user.Email,
		SignedUpAt: user.SignedUpAt,
	}}
	return unmarshalToUser(api.httpClient.Post("/visitors/convert", &cr))
}

func (api VisitorAPI) delete(id string) (Visitor, error) {
	visitor := Visitor{}
	data, err := api.httpClient.Delete(fmt.Sprintf("/visitors/%s", id), nil)
	if err != nil {
		return visitor, err
	}
	err = json.Unmarshal(data, &visitor)
	return visitor, err
}

type convertToLeadRequest struct {
	Lead    requestUser `json:"lead"`
	Visitor requestUser `json:"visitor"`
}

type convertToUserRequest struct {
	User    requestUser `json:"user"`
	Visitor requestUser `json:"visitor"`
}

func unmarshalToVisitor(data []byte, err error) (Visitor, error) {
	savedVisitor := Visitor{}
	if err != nil {
		return savedVisitor, err
	}
	err = json.Unmarshal(data, &savedVisitor)
	return savedVisitor, err
}

func (api VisitorAPI) buildRequestVisitor(visitor *Visitor) requestUser {
	return requestUser{
		ID:                     visitor.ID,
		Email:                  visitor.Email,
		Phone:                  visitor.Phone,
		UserID:                 visitor.UserID,
		Name:                   visitor.Name,
		LastRequestAt:          visitor.LastRequestAt,
		LastSeenIP:             visitor.LastSeenIP,
		UnsubscribedFromEmails: visitor.UnsubscribedFromEmails,
		Companies:              api.getCompaniesToSendFromVisitor(visitor),
		CustomAttributes:       visitor.CustomAttributes,
		UpdateLastRequestAt:    visitor.UpdateLastRequestAt,
		NewSession:             visitor.NewSession,
	}
}

func (api VisitorAPI) getCompaniesToSendFromVisitor(visitor *Visitor) []UserCompany {
	if visitor.Companies == nil {
		return []UserCompany{}
	}
	return RequestUserMapper{}.MakeUserCompaniesFromCompanies(visitor.Companies.Companies)
}
