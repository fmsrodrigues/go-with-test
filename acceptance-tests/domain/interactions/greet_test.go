package interactions

import (
	"acceptance-tests/specifications"
	"testing"
)

func TestGreet(t *testing.T) {
	specifications.GreetSpecification(t, specifications.GreetAdapter(Greet))
}
