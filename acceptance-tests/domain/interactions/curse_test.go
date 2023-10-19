package interactions_test

import (
	"acceptance-tests/domain/interactions"
	"acceptance-tests/specifications"
	"testing"
)

func TestCurse(t *testing.T) {
	specifications.CurseSpecification(t, specifications.CurseAdapter(interactions.Curse))
}
