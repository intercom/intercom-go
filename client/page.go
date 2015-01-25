package client

type PageParams struct {
	Page       int32 `json:"page" url:"page,omitempty"`
	PerPage    int32 `json:"per_page" url:"per_page,omitempty"`
	TotalPages int32 `json:"total_pages" url:"-"`
}
