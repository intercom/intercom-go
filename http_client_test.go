package intercom

type TestHTTPClient struct{}

func (h TestHTTPClient) Get(uri string, queryParams interface{}) ([]byte, error) { return nil, nil }
func (h TestHTTPClient) Post(uri string, body interface{}) ([]byte, error)       { return nil, nil }
func (h TestHTTPClient) Patch(uri string, body interface{}) ([]byte, error)      { return nil, nil }
func (h TestHTTPClient) Delete(uri string, body interface{}) ([]byte, error)     { return nil, nil }
