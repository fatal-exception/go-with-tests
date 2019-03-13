package iteration

import (
	"fmt"
	"testing"
)

func TestIt(t *testing.T) {
	t.Run("RepeatRecursive gives us a character five times", func(t *testing.T) {
		got := RepeatRecursive('q', 5)
		want := "qqqqq"
		if got != want {
			t.Errorf("want '%s', got '%s'", want, got)
		}
	})
	t.Run("Give5iterative gives us a character five times", func(t *testing.T) {
		got := RepeatIterative('q', 5)
		want := "qqqqq"
		if got != want {
			t.Errorf("want '%s', got '%s'", want, got)
		}
	})
	t.Run("Give5iterative takes number of time to be repeated", func(t *testing.T) {
		got := RepeatIterative('q', 6)
		want := "qqqqqq"
		if got != want {
			t.Errorf("want '%s', got '%s'", want, got)
		}
	})
}

func BenchmarkRepeatRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RepeatRecursive('b', 5)
	}
}

func BenchmarkRepeatIterative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RepeatIterative('b', 5)
	}
}

func ExampleRepeatIterative() {
	res := RepeatIterative('1', 7)
	fmt.Println(res)
	// Output: 1111111
}
