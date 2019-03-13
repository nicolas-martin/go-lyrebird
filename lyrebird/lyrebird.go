package lyrebird

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	authorizeURL   = "https://myvoice.lyrebird.ai/authorize"
	tokenURL       = "https://avatar.lyrebird.ai/api/v0/token"
	defaultBaseURL = "https://avatar.lyrebird.ai/api/v0/"
)

// Client manages communication with the Lyrebird API
type Client struct {
	client        *http.Client
	AvatarService *AvatarService
	VoiceService  *VoiceService
	UserAgent     string
}

type service struct {
	client *Client
}

// Response is a Lyrebird API response. This wraps the standard http.Response
// returned from Lyrebird and provides convenient access to things like
// pagination.
type Response struct {
	*http.Response

	// NextPage  int
	// PrevPage  int
	// FirstPage int
	// LastPage  int
}

// NewClient returns a new Lyrebird API client.
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	// baseURL, _ := url.Parse(defaultBaseURL)
	c := &Client{
		client: httpClient,
	}
	c.AvatarService = &AvatarService{client: c}
	c.VoiceService = &VoiceService{client: c}
	return c
}

// NewRequest does basic http request building
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	u, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	req.Proto = "HTTP/2"
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	// req.Header.Set("Accept", mediaTypeV3)
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}
	return req, nil

}

// Do sends API requests and returns API response
func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*Response, error) {
	// req = withContext(ctx, req)

	resp, err := c.client.Do(req)
	if err != nil {

		if e, ok := err.(*url.Error); ok {
			return nil, e
		}

		return nil, err
	}
	defer resp.Body.Close()
	response := &Response{Response: resp}
	b, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		return response, readErr
	}

	if response.Response.StatusCode != 200 {
		return response, errors.New(string(b))
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			decErr := json.Unmarshal(b, v)
			if decErr == io.EOF {
				decErr = nil // ignore EOF errors caused by empty response body
			}
			if decErr != nil {
				err = decErr
			}
		}
	}

	return response, err
}
