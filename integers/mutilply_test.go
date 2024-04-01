package integers

import (
	"testing"
)

func TestMultiply(t *testing.T) {
	product := Multiply(2, 3)
	expected := 6

	if product != expected {
		t.Errorf("expected '%d' but got '%d'", expected, product)
	}
}
