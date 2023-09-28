package adder_test

import (
	"testing"

	"github.com/banaaron/testing/adder"
)

func TestAdder(t *testing.T) {
	a, b := 1, 2
	result := adder.Add(a, b)
	if result != 3 {
		t.Errorf("1 + 2 does not equal %d", result)
	}
}
