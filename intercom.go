package intercom

import (
	"github.com/intercom/intercom-go/client"
	"github.com/intercom/intercom-go/interfaces"
)

type Client struct {
	User            client.User
	Event           client.Event
	userRepository  client.UserRepository
	eventRepository client.EventRepository
	httpClient      interfaces.HTTPClient
	baseURI         string
	debug           bool
}

const defaultBaseURI = "https://api.intercom.io"

func NewClient(appID, apiKey string) *Client {
	intercom := Client{baseURI: defaultBaseURI, debug: false}
	intercom.httpClient = interfaces.NewIntercomHTTPClient(appID, apiKey, &intercom.baseURI, &intercom.debug)
	intercom.userRepository = interfaces.NewUserAPI(intercom.httpClient)
	intercom.eventRepository = interfaces.NewEventAPI(intercom.httpClient)
	intercom.User = client.User{Repository: intercom.userRepository}
	intercom.Event = client.Event{Repository: intercom.eventRepository}
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
