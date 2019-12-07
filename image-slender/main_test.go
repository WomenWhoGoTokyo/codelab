package main

import "testing"

func BenchmarkExecute(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		execute()
	}
}
