package intercom

type Pages struct {
	Next       string `json:"next,omitempty"`
	Page       int32  `json:"page,omitempty"`
	PerPage    int32  `json:"per_page,omitempty"`
	TotalPages int32  `json:"total_pages,omitempty"`
}

type PageParams struct {
	Page    int32 `json:"page,omitempty" url:"page,omitempty"`
	PerPage int32 `json:"per_page,omitempty" url:"per_page,omitempty"`
}
