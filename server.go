package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("thing-registry server - starting...")
	log.Fatal(http.ListenAndServe(":4480", nil))
	log.Println("done - exiting.")
}
