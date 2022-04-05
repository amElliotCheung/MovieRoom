package main

import (
	"flag"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":8080", "http server address")

func main() {
	flag.Parse()
	http.HandleFunc("/ws", serveWs)
	log.Fatalln(http.ListenAndServe(*addr, nil))
}
