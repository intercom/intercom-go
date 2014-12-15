package intercom

type Note struct {
	*Resource
	User    userIdentifiers `json:"user"`
	AdminId string          `json:"admin_id,omitempty"`
	Body    string          `json:"body"`
}

type NoteParams struct {
	Id      string
	Email   string
	UserId  string
	AdminId string
	Body    string
}

func (n Note) New(params *NoteParams) (*Note, error) {
	note := Note{
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
