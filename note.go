package intercom

import "fmt"

type Note struct {
	*Resource
	Id     string          `json:"id"`
	User   userIdentifiers `json:"user"`
	Author Author          `json:"author,omitempty"`
	Body   string          `json:"body"`
}

type NoteParams struct {
	Id      string
	Email   string
	UserId  string
	AdminId string
	Body    string
}

type sentNote struct {
	User    userIdentifiers `json:"user"`
	AdminId string          `json:"admin_id,omitempty"`
	Body    string          `json:"body"`
}

func (n Note) New(params *NoteParams) (*Note, error) {
	note := sentNote{
		User: userIdentifiers{
			Id:     params.Id,
			Email:  params.Email,
			UserId: params.UserId,
		},
		AdminId: params.AdminId,
		Body:    params.Body,
	}
	if responseBody, err := n.client.Post("/notes", note); err != nil {
		return nil, err
	} else {
		newNote := Note{}
		return &newNote, n.Unmarshal(&newNote, responseBody.([]byte))
	}
}

func (n Note) Find(params *NoteParams) (*Note, error) {
	if responseBody, err := n.client.Get(fmt.Sprintf("/notes/%s", params.Id), nil); err != nil {
		return nil, err
	} else {
		newNote := Note{}
		return &newNote, n.Unmarshal(&newNote, responseBody.([]byte))
	}
}
