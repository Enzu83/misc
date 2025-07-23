package bench

import (
	"testing"
)

var num = 2000

func BenchmarkLoop(b *testing.B) {
	for b.Loop() {
		primeNumbers(num)
	}
}