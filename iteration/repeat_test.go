package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func TestRepeatWithCount(t *testing.T) {
	repeated := RepeatWithCount("x", 3)
	expected := "xxx"

	if repeated != expected {
		t.Errorf("expected %q to be %q", expected, repeated)
	}
}

func ExampleRepeat() {
	repeated := Repeat("b")
	fmt.Println(repeated)
	// Output: bbbbb
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}

func BenchmarkRepeatWithCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RepeatWithCount("x", 3)
	}
}
