package lyrebird

import (
	"context"
	"time"
)

const (
	generateURL  = "https://avatar.lyrebird.ai/api/v0/generate"
	generatedURL = "https://avatar.lyrebird.ai/api/v0/generated"
	profileURL   = "https://avatar.lyrebird.ai/api/v0/profile"
)

// AvatarService handles communication with the Avatar related
// methods of the Lyrebird API
type AvatarService service

// AvatarError represents am Lyrebird avatar error
type AvatarError struct {
	Description *string `json:"description,omitempty"`
}

// Avatar represents a Lyrebird avatar
type Avatar struct {
	CreatedAt *time.Time `json:"created_at,omitempty"`
	Text      *string    `json:"text,omitempty"`
	URL       *string    `json:"url,omitempty"`
}

func (a Avatar) String() string {
	return Stringify(a)
}

type generateBody struct {
	Text string
}

type Profile struct {
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
	UserID      string `json:"user_id"`
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

// Generated returns a list of generated voice clips
func (as *AvatarService) Generated(ctx context.Context) (*[]Avatar, *Response, error) {
	req, err := as.client.NewRequest("GET", generatedURL, nil)
	if err != nil {
		return nil, nil, err
	}

	avatar := new([]Avatar)
	resp, err := as.client.Do(ctx, req, avatar)
	if err != nil {
		return nil, resp, err
	}

	return avatar, resp, nil
}

// Profile returns the user's profile
func (as *AvatarService) Profile(ctx context.Context) (*Profile, *Response, error) {
	req, err := as.client.NewRequest("GET", profileURL, nil)
	if err != nil {
		return nil, nil, err
	}

	profile := new(Profile)
	resp, err := as.client.Do(ctx, req, profile)
	if err != nil {
		return nil, resp, err
	}

	return profile, resp, nil
}
