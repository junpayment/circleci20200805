package foo

import "testing"

func TestDo(t *testing.T) {
	if Do() != nil {
		t.Error("bar")
	}
}
