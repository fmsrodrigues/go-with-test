package main_test

import (
	"acceptance-tests/adapters"
	"acceptance-tests/adapters/grpcserver"
	"acceptance-tests/specifications"
	"fmt"
	"testing"
)

func TestGreeterServer(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	var (
		port   = "50051"
		driver = grpcserver.Driver{Addr: fmt.Sprintf("localhost:%s", port)}
	)

	adapters.StartDockerServer(t, port, "grpcserver")
	specifications.GreetSpecification(t, &driver)
}
