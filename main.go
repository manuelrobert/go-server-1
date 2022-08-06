package main

import (
	"log"
	"net/http"

	"github.com/manuelrobert/test_1_server/routes"
)

func main() {
	r := routes.Router()
	log.Fatal(http.ListenAndServe(":4000", r))
}
