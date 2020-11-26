package unittest

import "testing"

func TestAdd(t *testing.T) {
	if r := Add(1, 2); r != 3 {
		t.Errorf("1 + 2 expected 3, but %d", r)
	}
}
