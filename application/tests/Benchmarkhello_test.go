package tests

import (
		"testing"
		"fmt"		
		)

func Benchmark_Hello(b *testing.B) {
    for i := 0; i < b.N; i++ {
        fmt.Sprintf("hello")
    }
}