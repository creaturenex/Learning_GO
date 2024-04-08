package structs

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	actual := Perimeter(10.0, 10.0)
	expected := 40.0

	if actual != expected {
		t.Errorf("expected %.2f to be %.2f", expected, actual)
	}
}
