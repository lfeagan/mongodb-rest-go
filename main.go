package main

import (
	"log"
	"net/http"
	"flag"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"strconv"
	"fmt"
)

func main() {
	flag.String("listenHost", "", "the interface address to bind to for listening")
	flag.Int("listenPort", 8080, "the port number to listen for connections on")
	flag.String("mongoHost", "localhost", "the MongoDB host to connect to")
	flag.Int("mongoPort", 27017, "the MongoDB port to connect to")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	mongoHost := viper.GetString("mongoHost") // retrieve value from viper
	mongoPort := viper.GetInt("mongoPort")
	mongoAddr = mongoHost + ":" + strconv.Itoa(mongoPort)

	listenHost := viper.GetString("listenHost") // retrieve value from viper
	listenPort := viper.GetInt("listenPort")
	listenAddr := listenHost + ":" + strconv.Itoa(listenPort)
	fmt.Printf("listening on %s\n", listenAddr)
	router := NewRouter()
	log.Fatal(http.ListenAndServe(listenAddr, router))
}
