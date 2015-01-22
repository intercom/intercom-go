package interfaces

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type HTTPClient interface {
	Get(string) ([]byte, error)
	Post(string, interface{}) ([]byte, error)
}

const defaultBaseURI = "https://api.intercom.io"

type IntercomHttpClient struct {
	*http.Client
	BaseURI string
	AppId   string
	APIKey  string
	Debug   bool
}

func NewIntercomHTTPClient(appId, apiKey string) *IntercomHttpClient {
	return &IntercomHttpClient{Client: &http.Client{}, AppId: appId, APIKey: apiKey, BaseURI: defaultBaseURI, Debug: true}
}

func (c IntercomHttpClient) Get(url string) ([]byte, error) {
	req, _ := c.newRequest("GET", url, nil)
	req.SetBasicAuth(c.AppId, c.APIKey)
	req.Header.Add("Accept", "application/json")

	resp, _ := c.Client.Do(req)
	defer resp.Body.Close()

	return c.readAll(resp.Body)
}

func (c IntercomHttpClient) Post(url string, body interface{}) ([]byte, error) {
	buffer := bytes.NewBuffer([]byte{})
	if err := json.NewEncoder(buffer).Encode(body); err != nil {
		return nil, err
	}
	req, _ := c.newRequest("POST", url, buffer)

	req.SetBasicAuth(c.AppId, c.APIKey)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	resp, _ := c.Client.Do(req)
	defer resp.Body.Close()
	return c.readAll(resp.Body)
}

func (c IntercomHttpClient) newRequest(method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, c.BaseURI+url, body)
	if c.Debug {
		fmt.Printf("%s %s\n", req.Method, req.URL)
	}
	return req, err
}

func (c IntercomHttpClient) readAll(body io.Reader) ([]byte, error) {
	b, err := ioutil.ReadAll(body)
	if c.Debug {
		fmt.Println(string(b))
	}
	return b, err
}
