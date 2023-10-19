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
	if testing.Short() {
		t.Skip()
	}

	var (
		port   = "8080"
		driver = httpserver.Driver{
			BaseURL: fmt.Sprintf("http://localhost:%s", port),
			Client:  &http.Client{Timeout: 1 * time.Second},
		}
	)

	adapters.StartDockerServer(t, port, "httpserver")
	specifications.GreetSpecification(t, driver)
	specifications.CurseSpecification(t, driver)
}
