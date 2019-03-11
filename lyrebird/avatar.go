package lyrebird

import (
	"context"
)

const (
	generateURL = "https://avatar.lyrebird.ai/api/v0/generate"
)

// AvatarService handles communication with the Avatar related
// methods of the Lyrebird API
type AvatarService service

// Avatar represents a Lyrebird avatar
type Avatar struct {
	Description *string `json:"description,omitempty"`
}

func (a Avatar) String() string {
	return Stringify(a)
}

type generateBody struct {
	Text string
}

// Generate a sound clip from a string
func (as *AvatarService) Generate(ctx context.Context, text string) (*Avatar, *Response, error) {
	genBody := generateBody{
		Text: text,
	}
	req, err := as.client.NewRequest("POST", generateURL, genBody)
	if err != nil {
		return nil, nil, err
	}

	avatar := new(Avatar)
	resp, err := as.client.Do(ctx, req, avatar)
	if err != nil {
		return nil, resp, err
	}

	return avatar, resp, nil
}
