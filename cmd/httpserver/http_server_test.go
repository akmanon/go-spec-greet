package main_test

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/akmanon/go-spec-greet/adapters"
	"github.com/akmanon/go-spec-greet/adapters/httpserver"
	"github.com/akmanon/go-spec-greet/specifications"
)

func TestHttpServer(t *testing.T) {
	var (
		port           = "8080"
		dockerFilePath = "./cmd/httpserver/Dockerfile"
		baseURL        = fmt.Sprintf("http://localhost:%s", port)
		driver         = httpserver.Driver{BaseURL: baseURL, Client: &http.Client{
			Timeout: 1 * time.Second,
		}}
	)
	adapters.StartDockerServer(t, port, dockerFilePath)
	t.Run("Greet Specification", func(t *testing.T) {
		specifications.GreetSpecification(t, driver)
	})
	t.Run("Curse Specification", func(t *testing.T) {
		specifications.CurseSpecification(t, driver)
	})

}
