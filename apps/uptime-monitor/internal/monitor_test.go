package monitor

import "testing"

// TestGet checks that Get returns one of the strings from fortunes.
func TestAdd(t *testing.T) {
	res := add(1, 2)
	if i := res; i != 3 {
		t.Errorf("add returned %q, not 3", res)
	}
}
