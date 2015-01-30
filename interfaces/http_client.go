package interfaces

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/google/go-querystring/query"
	"github.com/intercom/intercom-go/client"
)

type HTTPClient interface {
	Get(string, interface{}) ([]byte, error)
	Post(string, interface{}) ([]byte, error)
}

type IntercomHttpClient struct {
	*http.Client
	BaseURI *string
	AppId   string
	APIKey  string
	Debug   *bool
}

func NewIntercomHTTPClient(appId, apiKey string, baseURI *string, debug *bool) IntercomHttpClient {
	return IntercomHttpClient{Client: &http.Client{}, AppId: appId, APIKey: apiKey, BaseURI: baseURI, Debug: debug}
}

func (c IntercomHttpClient) Get(url string, queryParams interface{}) ([]byte, error) {
	// Setup request
	req, _ := http.NewRequest("GET", *c.BaseURI+url, nil)
	req.SetBasicAuth(c.AppId, c.APIKey)
	req.Header.Add("Accept", "application/json")
	addQueryParams(req, queryParams)
	if *c.Debug {
		fmt.Printf("%s %s\n", req.Method, req.URL)
	}

	// Do request
	resp, err := c.Client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	// Read response
	data, err := c.readAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 400 {
		return nil, c.parseResponseError(data, resp.StatusCode)
	}
	return data, err
}

func addQueryParams(req *http.Request, params interface{}) {
	v, _ := query.Values(params)
	req.URL.RawQuery = v.Encode()
}

func (c IntercomHttpClient) Post(url string, body interface{}) ([]byte, error) {
	// Marshal our body
	buffer := bytes.NewBuffer([]byte{})
	if err := json.NewEncoder(buffer).Encode(body); err != nil {
		return nil, err
	}

	// Setup request
	req, err := http.NewRequest("POST", *c.BaseURI+url, buffer)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(c.AppId, c.APIKey)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if *c.Debug {
		fmt.Printf("%s %s\n", req.Method, req.URL)
	}

	// Do request
	resp, err := c.Client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	// Read response
	data, err := c.readAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 400 {
		return nil, c.parseResponseError(data, resp.StatusCode)
	}
	return data, err
}

func (c IntercomHttpClient) parseResponseError(data []byte, statusCode int) client.IntercomError {
	errorList := HTTPErrorList{}
	err := json.Unmarshal(data, &errorList)
	if err != nil {
		return NewUnknownHTTPError()
	}
	httpError := errorList.Errors[0]
	httpError.StatusCode = statusCode
	return httpError // only care about the first
}

func (c IntercomHttpClient) readAll(body io.Reader) ([]byte, error) {
	b, err := ioutil.ReadAll(body)
	if *c.Debug {
		fmt.Println(string(b))
	}
	return b, err
}
