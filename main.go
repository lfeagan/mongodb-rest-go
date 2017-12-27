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
	flag.String("hostName", "", "the interface address to bind to for listening")
	flag.Int("hostPort", 8080, "the port to listen on for connections")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	hostName := viper.GetString("hostName") // retrieve value from viper
	hostPort := viper.GetInt("hostPort")

	listenAddr := hostName + ":" + strconv.Itoa(hostPort)
	fmt.Printf("listening on %s\n", listenAddr)
	router := NewRouter()
	log.Fatal(http.ListenAndServe(listenAddr, router))
}
