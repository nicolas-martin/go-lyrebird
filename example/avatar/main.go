package main

import (
	"context"
	"fmt"
	"log"
	"lyrebird-go/lyrebird"
	"net/url"
	"os"

	"golang.org/x/oauth2"
)

func main() {
	lyrebirdCode := os.Getenv("LYREBIRD_CODE")
	ctx := context.Background()
	if len(lyrebirdCode) == 0 {

		lyrebirdClientID := os.Getenv("LYREBIRD_CLIENT_ID")
		lyrebirdSecret := os.Getenv("LYREBIRD_CLIENT_SECRET")

		conf := &oauth2.Config{
			ClientID:     lyrebirdClientID,
			ClientSecret: lyrebirdSecret,
			Scopes:       []string{"voice"},
			Endpoint: oauth2.Endpoint{
				AuthURL:  "https://myvoice.lyrebird.ai/authorize",
				TokenURL: "https://avatar.lyrebird.ai/api/v0/token",
			},
			RedirectURL: "http://aaaaaaabb.com/callback",
		}
		authurl := conf.AuthCodeURL("state", oauth2.AccessTypeOnline)

		// response_type=code doesn't work
		u0, err := url.Parse(authurl)
		if err != nil {
			fmt.Println("invalid url")
		}
		urlQuery := u0.Query()
		urlQuery.Set("response_type", "token")
		u0.RawQuery = urlQuery.Encode()
		fmt.Printf("Visit the URL for the auth dialog: %v \r\n", u0)

		if _, err := fmt.Scan(&lyrebirdCode); err != nil {
			log.Fatal(err)
		}
	}
	// tok, err := conf.Exchange(ctx, code)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// client := conf.Client(ctx, tok)
	client := oauth2.NewClient(ctx, oauth2.StaticTokenSource(&oauth2.Token{
		AccessToken: lyrebirdCode,
		TokenType:   "Bearer",
	}))

	lyreBirdClient := lyrebird.NewClient(client)

	a, r, err := lyreBirdClient.AvatarService.Generate(ctx, "generate something")
	if err != nil {
		fmt.Printf("error: %s \r\n", err.Error())
		fmt.Println(r.StatusCode)
		return
	}

	fmt.Println(a.Description)

}
