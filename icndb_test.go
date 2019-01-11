package icndb

import (
	"testing"
)

func TestPrepNames_success_timeSpace(t *testing.T) {
	first := "  chase"
	last := "isabelle   "

	names := prepNames(first, last)

	first = names["firstName"]
	last = names["lastName"]

	if first != "chase" {
		t.Errorf("Expected chase, but got %+v.", first)
	}

	if last != "isabelle" {
		t.Errorf("Expected isabelle, but got %+v.", last)
	}
}
