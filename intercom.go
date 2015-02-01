package intercom

import (
	"github.com/intercom/intercom-go/interfaces"
)

type Client struct {
	Admins                 AdminService
	Conversations          ConversationService
	Events                 EventService
	Users                  UserService
	AdminRepository        AdminRepository
	ConversationRepository ConversationRepository
	EventRepository        EventRepository
	UserRepository         UserRepository
	AppID                  string
	APIKey                 string
	HTTPClient             interfaces.HTTPClient
	baseURI                string
	debug                  bool
}

const defaultBaseURI = "https://api.intercom.io"

func NewClient(appID, apiKey string) *Client {
	intercom := Client{AppID: appID, APIKey: apiKey, baseURI: defaultBaseURI, debug: false}
	intercom.HTTPClient = interfaces.NewIntercomHTTPClient(intercom.AppID, intercom.APIKey, &intercom.baseURI, &intercom.debug)
	intercom.setup()
	return &intercom
}

func (c *Client) setup() {
	c.AdminRepository = AdminAPI{httpClient: c.HTTPClient}
	c.ConversationRepository = ConversationAPI{httpClient: c.HTTPClient}
	c.EventRepository = EventAPI{httpClient: c.HTTPClient}
	c.UserRepository = UserAPI{httpClient: c.HTTPClient}
	c.Admins = AdminService{Repository: c.AdminRepository}
	c.Conversations = ConversationService{Repository: c.ConversationRepository}
	c.Events = EventService{Repository: c.EventRepository}
	c.Users = UserService{Repository: c.UserRepository}
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

func SetHTTPClient(httpClient interfaces.HTTPClient) option {
	return func(c *Client) option {
		previous := c.HTTPClient
		c.HTTPClient = httpClient
		c.setup()
		return SetHTTPClient(previous)
	}
}
