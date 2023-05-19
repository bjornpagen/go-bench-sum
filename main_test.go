package main

import "testing"

// func Sum(fileName string) (ret int64, _ error) {
func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum("data.txt")
	}
}

func BenchmarkSumStream(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumStream("data.txt")
	}
}

func BenchmarkSumByteScanner(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumByteScanner("data.txt")
	}
}
