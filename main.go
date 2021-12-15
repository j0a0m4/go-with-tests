package main

import (
	"github.com/j0a0m4/go-with-tests/server"
	"log"
	"net/http"
)

const port = ":5050"

func main() {
	svr := &server.PlayerServer{Store: server.NewInMemoryPlayerStore()}
	log.Fatal(http.ListenAndServe(port, svr))
}
