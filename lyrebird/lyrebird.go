package lyrebird

import (
	"net/http"
	"net/url"
)

const (
	authorizeURL   = "https://myvoice.lyrebird.ai/authorize"
	generateURL    = "https://avatar.lyrebird.ai/api/v0/generate"
	tokenURL       = "https://avatar.lyrebird.ai/api/v0/token"
	defaultBaseURL = "https://avatar.lyrebird.ai/api/v0/"
)

// Client manages communication with the Lyrebird API
type Client struct {
	client        *http.Client
	BaseURL       *url.URL
	AvatarService *AvatarService
	VoiceService  *VoiceService
}

type service struct {
	client *Client
}

// NewClient returns a new Lyrebird API client.
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	baseURL, _ := url.Parse(defaultBaseURL)
	c := &Client{
		client:  httpClient,
		BaseURL: baseURL,
	}
	c.AvatarService = &AvatarService{client: c}
	c.VoiceService = &VoiceService{client: c}
	return c
}
