package intercom

import (
	"github.com/intercom/intercom-go/interfaces"
)

type Client struct {
	Users           UserService
	Events          EventService
	UserRepository  UserRepository
	EventRepository EventRepository
	HTTPClient      interfaces.HTTPClient
	baseURI         string
	debug           bool
}

const defaultBaseURI = "https://api.intercom.io"

func NewClient(appID, apiKey string) *Client {
	intercom := Client{baseURI: defaultBaseURI, debug: false}
	intercom.HTTPClient = interfaces.NewIntercomHTTPClient(appID, apiKey, &intercom.baseURI, &intercom.debug)
	intercom.UserRepository = UserAPI{httpClient: intercom.HTTPClient}
	intercom.EventRepository = EventAPI{httpClient: intercom.HTTPClient}
	intercom.Users = UserService{Repository: intercom.UserRepository}
	intercom.Events = EventService{Repository: intercom.EventRepository}
	return &intercom
}

type option func(c *Client) option

func (c *Client) Option(opts ...option) (previous option) {
	for _, opt := range opts {
		previous = opt(c)
	}
	return previous
}

func TraceHTTP(trace bool) option {
	return func(c *Client) option {
		previous := c.debug
		c.debug = trace
		return TraceHTTP(previous)
	}
}

func BaseURI(baseURI string) option {
	return func(c *Client) option {
		previous := c.baseURI
		c.baseURI = baseURI
		return BaseURI(previous)
	}
}
