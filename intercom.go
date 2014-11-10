// Package intercom provides bindings for Intercom API's
package intercom

import (
	"encoding/json"
	"log"

	"github.com/franela/goreq"
)

const defaultBaseUri = "https://api.intercom.io"
const clientVersion = "0.0.1"

type Client struct {
	AppId   string
	ApiKey  string
	BaseUri string
	Events  *Event
	trace   bool
}

func GetClient(appId string, apiKey string) *Client {
	c := Client{AppId: appId, ApiKey: apiKey, BaseUri: defaultBaseUri}
	c.Events = (&Event{}).SetClient(&c)
	return &c
}

type option func(c *Client) option

func (c *Client) Option(opts ...option) (previous option) {
	for _, opt := range opts {
		previous = opt(c)
	}
	return previous
}

func TraceHttp(trace bool) option {
	return func(c *Client) option {
		previous := c.trace
		c.trace = trace
		return TraceHttp(previous)
	}
}

func BaseUri(baseUri string) option {
	return func(c *Client) option {
		previous := c.BaseUri
		c.BaseUri = baseUri
		return BaseUri(previous)
	}
}

func (c *Client) Post(url string, body interface{}) (*goreq.Response, error) {
	return c.withTracing(c.postRequestTracer(), c.raisingPoster(), url, body)
}

func (c *Client) withTracing(
	request_tracer func(string, interface{}),
	http_function func(*Client, string, interface{}) (*goreq.Response, error),
	url string, body interface{}) (*goreq.Response, error) {

	if c.trace {
		request_tracer(c.BaseUri+url, body)
	}
	res, err := http_function(c, url, body)
	if c.trace {
		c.responseTracer()(res, err)
	}
	return res, err
}

func (c *Client) postRequestTracer() func(string, interface{}) {
	return func(url string, body interface{}) {
		log.Printf("[intercom] POSTing request to %s", url)
		b, err := json.Marshal(body)
		if err == nil {
			log.Printf("[intercom] with body %s", b)
		}
	}
}

func (c *Client) responseTracer() func(*goreq.Response, error) {
	return func(res *goreq.Response, httpError error) {
		if httpError != nil {
			log.Printf("[intercom] returned with error: %s", httpError)
		}
		if res == nil {
			return
		}
		log.Printf("[intercom] returned status: %d", res.StatusCode)
		bodyString, err := res.Body.ToString()
		if err == nil {
			log.Printf("[intercom] with body: %s", bodyString)
		}
	}
}

func (c *Client) raisingPoster() func(*Client, string, interface{}) (*goreq.Response, error) {
	return func(c *Client, url string, body interface{}) (*goreq.Response, error) {
		res, err := c.poster()(c, url, body)
		if err != nil {
			return nil, err // return if we've an unexpected error
		}
		switch res.StatusCode { // deal with error status codes
		case 400:
			errorList := HttpErrorList{}
			e := res.Body.FromJsonTo(&errorList)
			if e != nil {
				return nil, NewUnknownHttpError(res.StatusCode)
			}
			httpError := errorList.Errors[0] // only care about the first
			httpError.StatusCode = res.StatusCode
			return nil, &httpError
		}
		return res, nil
	}
}

func (c *Client) poster() func(*Client, string, interface{}) (*goreq.Response, error) {
	return func(c *Client, url string, body interface{}) (*goreq.Response, error) {
		return goreq.Request{
			Method:            "POST",
			Uri:               c.BaseUri + url,
			Accept:            "application/json",
			ContentType:       "application/json",
			BasicAuthUsername: c.AppId,
			BasicAuthPassword: c.ApiKey,
			Body:              body,
			UserAgent:         "intercom-go/" + clientVersion,
		}.Do()
	}
}
