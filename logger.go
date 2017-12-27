package main

import (
	"log"
	"net/http"
	"time"
	"net/url"
)

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			prettyUrl(r),
			name,
			time.Since(start),
		)
	})
}

func prettyUrl(r *http.Request) string {
	pretty,err := url.QueryUnescape(r.RequestURI)
	if err != nil {
		panic(err)
	}
	return pretty
}
