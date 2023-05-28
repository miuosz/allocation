package allocation

import (
	"testing"
)

func Benchmark_New(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		New(1, Byte, false, nil)
	}
}
