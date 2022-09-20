package iteration

import (
	"fmt"
	"testing"
)

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		repeated := Repeat("a")
		expected := "aaaaa"

		if repeated != expected {
			b.Errorf("Expected %q, got %q", expected, repeated)
		}
	}
}

func ExampleRepeat() {
	repeat := Repeat("a")
	fmt.Println(repeat)
	// output: aaaaa
}
