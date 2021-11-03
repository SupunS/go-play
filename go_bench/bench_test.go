package go_bench

import (
	"testing"
	"time"
)

func BenchmarkSomething(b *testing.B) {
	for i := 0; i < b.N; i++ {
		time.Sleep(200)
	}
}
