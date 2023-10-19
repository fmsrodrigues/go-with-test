package interactions_test

import (
	"acceptance-tests/domain/interactions"
	"acceptance-tests/specifications"
	"testing"
)

func TestGreet(t *testing.T) {
	specifications.GreetSpecification(t, specifications.GreetAdapter(interactions.Greet))
}
