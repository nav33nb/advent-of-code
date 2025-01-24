package day4

import (
	"testing"
)

func BenchmarkD4P1(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		Puz1()
	}
}
