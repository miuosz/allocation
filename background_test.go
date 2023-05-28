package allocation

import "testing"

func Benchmark_NewBackground(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewBackground(1, Byte, false, nil)
	}
}
