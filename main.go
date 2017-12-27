package main

import (
	"log"
	"net/http"
)

func main() {
	hostName := "localhost"
	hostPort := "8080"
	listenAddr := hostName + ":" + hostPort
	router := NewRouter()
	log.Fatal(http.ListenAndServe(listenAddr, router))
}
