package iteration

import (
	"strings"
	"testing"
)

func TestCount(t *testing.T) {
	actual := strings.Count("Hello, World", "")
	expected := 12

	if actual != expected {
		t.Errorf("expected %d to be %d", actual, expected)
	}
}
func TestCount1arg(t *testing.T) {
	actual := strings.Count("Hello, World")
	expected := 13

	if actual != expected {
		t.Errorf("expected %d to be %d", actual, expected)
	}
}
