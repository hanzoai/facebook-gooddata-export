package main

import (
	"fmt"
	"net/http"
)

func oauthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func oauthRedirectHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/redirect/", oauthRedirectHandler)
	http.ListenAndServe(":8080", nil)
}
