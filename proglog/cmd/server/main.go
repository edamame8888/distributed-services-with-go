package main

import (
	"log"

	"github.com/edamame8888/distributed-services-with-go/proglog/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")

	log.Fatal(srv.ListenAndServe())
}
