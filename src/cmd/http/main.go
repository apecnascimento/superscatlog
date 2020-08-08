package main

import (
	"flag"
	"supercatlog/src/api"
)

var (
	httpServerAddr = flag.String("http-addr", ":8080", "TCP address that the http server will listen on")
)

func main() {
	flag.Parse()
	api := api.Routes()
	api.Logger.Fatal(api.Start(*httpServerAddr))
}
