package intercom

type Resource struct {
	client *Client
}

func (r *Resource) SetClient(client *Client) *Resource {
	r.client = client
	return r
}
