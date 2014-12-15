package intercom

import "encoding/json"

type Resource struct {
	client *Client
}

func (r *Resource) SetClient(client *Client) *Resource {
	r.client = client
	return r
}

func (r *Resource) Unmarshal(o interface{}, responseBody []byte) error {
	err := json.Unmarshal(responseBody, &o)
	return err
}
