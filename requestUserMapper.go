package intercom

// A Company the User belongs to
// used to update Companies on a User.
type UserCompany struct {
	CompanyID        string                 `json:"company_id,omitempty"`
	Name             string                 `json:"name,omitempty"`
	Remove           *bool                  `json:"remove,omitempty"`
	CustomAttributes map[string]interface{} `json:"custom_attributes,omitempty"`
}

type RequestUserMapper struct{}

func (rum RequestUserMapper) ConvertUser(user *User) requestUser {
	return requestUser{
		ID:                     user.ID,
		Email:                  user.Email,
		Phone:                  user.Phone,
		UserID:                 user.UserID,
		Name:                   user.Name,
		SignedUpAt:             user.SignedUpAt,
		RemoteCreatedAt:        user.RemoteCreatedAt,
		LastRequestAt:          user.LastRequestAt,
		LastSeenIP:             user.LastSeenIP,
		UnsubscribedFromEmails: user.UnsubscribedFromEmails,
		Companies:              rum.getCompaniesToSendFromUser(user),
		CustomAttributes:       user.CustomAttributes,
		UpdateLastRequestAt:    user.UpdateLastRequestAt,
		NewSession:             user.NewSession,
		LastSeenUserAgent:      user.LastSeenUserAgent,
	}
}

func (rum RequestUserMapper) getCompaniesToSendFromUser(user *User) []UserCompany {
	if user.Companies == nil {
		return []UserCompany{}
	}
	return rum.MakeUserCompaniesFromCompanies(user.Companies.Companies)
}

func (rum RequestUserMapper) MakeUserCompaniesFromCompanies(companies []Company) []UserCompany {
	userCompanies := make([]UserCompany, len(companies))
	for i := 0; i < len(companies); i++ {
		userCompanies[i] = UserCompany{
			CompanyID:        companies[i].CompanyID,
			Name:             companies[i].Name,
			Remove:           companies[i].Remove,
			CustomAttributes: companies[i].CustomAttributes,
		}
	}
	return userCompanies
}
