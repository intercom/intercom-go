package intercom

import "time"

type Event struct {
	*Resource
	UserId    string       `json:"user_id,omitempty"`
	Email     string       `json:"email,omitempty"`
	Id        string       `json:"id,omitempty"`
	EventName string       `json:"event_name"`
	CreatedAt int32        `json:"created_at"`
	Metadata  AttributeMap `json:"metadata,omitempty"`
}

type EventParams struct {
	UserId    string
	Email     string
	Id        string
	EventName string
	CreatedAt int32
	Metadata  AttributeMap
}

func (e Event) New(params EventParams) error {
	event := Event{
		UserId:    params.UserId,
		Email:     params.Email,
		Id:        params.Id,
		EventName: params.EventName,
		CreatedAt: params.CreatedAt,
		Metadata:  params.Metadata,
	}
	if params.CreatedAt <= 0 {
		event.CreatedAt = int32(time.Now().Unix())
	}
	_, err := e.client.Post("/events", event)
	return err
}
