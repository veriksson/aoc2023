package util

import "testing"

func TestInput(t *testing.T) {
	lines := Input("test.txt")
	if len(lines) != 3 {
		t.FailNow()
	}
}
