package main

import (
	"log"
	"net/http"

	go_spec_greet "github.com/akmanon/go-spec-greet"
)

func main() {
	handler := http.HandlerFunc(go_spec_greet.Handler)
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}
