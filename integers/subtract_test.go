package integers

import "testing"

func TestSubtract(t *testing.T) {
	diff := Subtract(5, 3)
	expected := 2

	if diff != expected {
		t.Errorf("expect '%d' to be '%d'", expected, diff)
	}
}
