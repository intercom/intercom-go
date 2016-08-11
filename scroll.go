package intercom

type ScrollParams struct {
	Param string `json:"scroll_param" url:"scroll_param,omitempty"`
	Type  string `json:"type" url:"type,omitempty"`
}
