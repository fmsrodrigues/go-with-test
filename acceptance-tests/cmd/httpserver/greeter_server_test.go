package main_test

import (
	"acceptance-tests/adapters"
	"acceptance-tests/adapters/httpserver"
	"acceptance-tests/specifications"
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestGreeterServer(t *testing.T) {
	var (
		port           = "8080"
		dockerFilePath = "./Dockerfile"
		baseURL        = fmt.Sprintf("http://localhost:%s", port)
		driver         = httpserver.Driver{BaseURL: baseURL, Client: &http.Client{Timeout: 1 * time.Second}}
	)

	adapters.StartDockerServer(t, port, dockerFilePath)
	specifications.GreetSpecification(t, driver)
}
