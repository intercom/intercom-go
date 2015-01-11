// Package intercom provides bindings for Intercom API's
package intercom

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/franela/goreq"
	"github.com/google/go-querystring/query"
)

const defaultBaseUri = "https://api.intercom.io"
const clientVersion = "0.0.1"

type Client struct {
	AppId   string
	ApiKey  string
	BaseUri string
	Events  *Event
	Users   *User
	Notes   *Note
	Admins  *Admin
	trace   bool
}

func GetClient(appId string, apiKey string) *Client {
	c := Client{AppId: appId, ApiKey: apiKey, BaseUri: defaultBaseUri}
	resource := (&Resource{}).SetClient(&c)
	c.Events = &Event{Resource: resource}
	c.Notes = &Note{Resource: resource}
	c.Users = &User{Resource: resource}
	c.Admins = &Admin{Resource: resource}
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

func (c *Client) Post(url string, body interface{}) (interface{}, error) {
	return c.withTracing(c.postRequestTracer(), c.raisingPoster(), url, body)
}

func (c *Client) Get(url string, queryObject interface{}) (interface{}, error) {
	return c.withTracing(c.getRequestTracer(), c.raisingGetter(), url, queryObject)
}

func (c *Client) withTracing(
	request_tracer func(string, interface{}),
	http_function func(*Client, string, interface{}) (*goreq.Response, error),
	url string, body interface{}) (interface{}, error) {

	if c.trace {
		request_tracer(c.BaseUri+url, body)
	}
	res, err := http_function(c, url, body)
	bodyBytes, err := ioutil.ReadAll(res.Body)
	if c.trace {
		if res != nil && err == nil {
			c.responseTracerBody()(res.StatusCode, string(bodyBytes), err)
		}
	}
	return bodyBytes, err
}

func (c *Client) postRequestTracer() func(string, interface{}) {
	return func(url string, body interface{}) {
		log.Printf("[intercom] POST request to %s", url)
		b, err := json.Marshal(body)
		if err == nil {
			log.Printf("[intercom] with body %s", b)
		}
	}
}

func (c *Client) getRequestTracer() func(string, interface{}) {
	return func(url string, queryObject interface{}) {
		v, err := query.Values(queryObject)
		if err == nil && len(v) != 0 {
			log.Printf("[intercom] GET request to %s?%s", url, v.Encode())
		} else if err == nil {
			log.Printf("[intercom] GET request to %s", url)
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

func (c *Client) responseTracerBody() func(int, string, error) {
	return func(statusCode int, bodyString string, httpError error) {
		if httpError != nil {
			log.Printf("[intercom] returned with error: %s", httpError)
		}
		log.Printf("[intercom] returned status: %d", statusCode)
		if bodyString != "" {
			log.Printf("[intercom] with body: %s", bodyString)
		}
	}
}

func (c *Client) raisingPoster() func(*Client, string, interface{}) (*goreq.Response, error) {
	return func(c *Client, url string, body interface{}) (*goreq.Response, error) {
		res, err := c.poster()(c, url, body)
		// fmt.Printf(res.Body.ToString())
		return c.parseResult(res, err)
	}
}

func (c *Client) raisingGetter() func(*Client, string, interface{}) (*goreq.Response, error) {
	return func(c *Client, url string, queryObject interface{}) (*goreq.Response, error) {
		res, err := c.getter()(c, url, queryObject)
		return c.parseResult(res, err)
	}
}

func (c *Client) parseResult(res *goreq.Response, err error) (*goreq.Response, error) {
	if err != nil {
		return nil, err // return if we've an unexpected error
	}
	if res.StatusCode >= 400 {
		errorList := HttpErrorList{}
		e := res.Body.FromJsonTo(&errorList)
		if e != nil {
			return nil, NewUnknownHttpError(res.StatusCode)
		}
		httpError := errorList.Errors[0] // only care about the first
		httpError.StatusCode = res.StatusCode
		return res, &httpError
	}
	return res, nil
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

func (c *Client) getter() func(*Client, string, interface{}) (*goreq.Response, error) {
	return func(c *Client, url string, queryObject interface{}) (*goreq.Response, error) {
		queryString, _ := query.Values(queryObject)
		return goreq.Request{
			Method:            "GET",
			Uri:               c.BaseUri + url,
			Accept:            "application/json",
			ContentType:       "application/json",
			BasicAuthUsername: c.AppId,
			BasicAuthPassword: c.ApiKey,
			QueryString:       queryString,
			UserAgent:         "intercom-go/" + clientVersion,
		}.Do()
	}
}
