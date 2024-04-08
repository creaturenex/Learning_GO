package iteration

import (
	"strings"
	"testing"
)

func TestCut(t *testing.T) {
	before, after, found := strings.Cut("Hello, World", ",")
	expectedBefore, expectedAfter, expectedBoolean := "Hello", " World", true

	if before != expectedBefore && after != expectedAfter {
		t.Errorf("expected %q to be %q", before, expectedBefore)
	}
	if after != expectedAfter {
		t.Errorf("expected %q to be %q", after, expectedAfter)
	}
	if found != expectedBoolean {
		t.Errorf("expected %v to be %v", found, expectedBoolean)
	}
}
