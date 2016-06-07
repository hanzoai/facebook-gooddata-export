package main

import (
	"errors"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/flosch/pongo2"
	"github.com/hanzo-io/oauthful"
	"github.com/skratchdot/open-golang/open"
)

var redirectHTML = pongo2.Must(pongo2.FromFile("templates/redirect.html"))

type FacebookFlow struct {
}

func (f FacebookFlow) Decode(req *http.Request) (*oauthful.AuthorizationResponse, error) {
	errStr := req.FormValue("error")
	if errStr != "" {
		return nil, errors.New(errStr + "\n" + req.FormValue("error_reason") + "\n" + req.FormValue("error_description"))
	}

	res := &oauthful.AuthorizationResponse{
		State: req.FormValue("state"),
		Code:  req.FormValue("code"),
	}

	return res, nil
}

func (f FacebookFlow) Verify(res *oauthful.AuthorizationResponse) error {
	if res.State != Config.State {
		return errors.New("Authorization Error, State Does Not Match")
	}
	return nil
}

func (f FacebookFlow) AddParams(vals *url.Values) error {
	vals.Add("client_id", Config.ClientId)
	vals.Add("client_secret", Config.ClientSecret)
	vals.Add("redirect_uri", Config.RedirectUri)
	return nil
}

func oauthRedirectHandler(w http.ResponseWriter, r *http.Request) {
	hc := &http.Client{}
	client := oauthful.New(hc, Config.TokenUrl, FacebookFlow{})

	pctx := pongo2.Context{
		"success": true,
	}

	if res, err := client.Handle(r); err != nil {
		pctx["success"] = false
		pctx["error"] = err.Error()
	} else {
		pctx["response"] = res
	}

	err := redirectHTML.ExecuteWriter(pctx, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	go server()

	qs := []string{
		"client_id=" + Config.ClientId,
		"state=" + Config.State,
		"scope=" + Config.Scope,
		"redirect_uri=" + Config.RedirectUri,
	}

	url := Config.AuthorizeUrl + "?" + strings.Join(qs, "&")

	open.Run(url)

	time.Sleep(time.Second * 10)
}

func server() {
	http.HandleFunc("/redirect", oauthRedirectHandler)
	http.ListenAndServe(":8080", nil)
}
