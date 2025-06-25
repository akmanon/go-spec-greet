package main

import (
	"log"
	"net/http"

	go_spec_greet "github.com/akmanon/go-spec-greet"
)

func main() {
	handlerGreet := http.HandlerFunc(go_spec_greet.GreetHandler)
	handlerCurse := http.HandlerFunc(go_spec_greet.CurseHandler)
	http.HandleFunc("/greet", handlerGreet)
	http.HandleFunc("/curse", handlerCurse)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
