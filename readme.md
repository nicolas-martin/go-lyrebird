# go-lyrebird #

go-lyrebird is a Go client library for accessing the lyrebird API.

## Usage ##

To create a new Lyrebird client and use it's services. Create a new `*http.Client` with the required authentication token. Go-lyrebird does not directly handle authentication, instead it is suggested to use the [oauth2](https://github.com/golang/oauth2) library.

``` GO
ctx := context.Background()
client := oauth2.NewClient(ctx, oauth2.StaticTokenSource(&oauth2.Token{
    AccessToken: lyrebirdCode,
    TokenType:   "Bearer",
})

lyreBirdClient := lyrebird.NewClient(client)

p, r, err := lyreBirdClient.AvatarService.Profile(ctx)
if err != nil {
    log.Fatal(err)
}

fmt.Println(p.DisplayName)
```

## Getting started with the examples ##

- If you do not have an api key, you will need to set `LYREBIRD_CLIENT_ID`, `LYREBIRD_CLIENT_SECRET` which can be found in the lyrebird application settings page anda re required in order obtain an api key.

- If you already have a key avaiable, you can store it in the `LYREBIRD_CODE` which bypass the OAuth2 workflow.
