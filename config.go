package main

type config struct {
	State string

	AuthorizeUrl string
	TokenUrl     string
	RedirectUri  string

	ClientId     string
	ClientSecret string
	Scope        string
}

var Config = config{
	State: "12345678",

	AuthorizeUrl: "https://www.facebook.com/dialog/oauth",
	TokenUrl:     "https://graph.facebook.com/v2.3/oauth/access_token",
	RedirectUri:  "http://localhost:8080/redirect",

	ClientId:     "244032362636714",
	ClientSecret: "b680c55a6ef62e603feaf1f33c0e15f1",
	Scope:        "read_insights,ads_read",
}
