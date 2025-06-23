package main_test

import (
	"testing"

	go_specs_greet "github.com/akmanon/go-spec-greet"
	"github.com/akmanon/go-spec-greet/specifications"
)

func TestGreeterServer(t *testing.T) {
	driver := go_specs_greet.Driver{BaseURL: "http://localhost:8080"}
	specifications.GreetSpecification(t, nil)
}
