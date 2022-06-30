package main

import (
	"log"
	"time"
	"net/http"
	"net/url"
	"strconv"
)

func main() {
	s := &http.Server{
		Addr:           ":8080",
		Handler:        http.HandlerFunc(testHandle),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}

func testHandle(w http.ResponseWriter, r *http.Request) {
	log.Println(fmtURL(url.ParseRequestURI(r.URL))
}

func fmtURL(u *url.URL) string {
	res := "\n"
	res += "Scheme: " + u.Scheme + "\n"
	res += "Opaque: " + u.Opaque + "\n"
	res += "Host: " + u.Hostname() + "\n"
	res += "Path: " + u.Path + "\n"
	res += "RawPath: " + u.RawPath + "\n"
	res += "ForceQuery: " + strconv.FormatBool(u.ForceQuery) + "\n"
	res += "RawQuery: " + u.RawQuery + "\n"
	res += "Fragment: " + u.Fragment + "\n"
	res += "RawFragment: " + u.RawFragment + "\n"
	return res
}