package main

import (
	"math"
	"testing"
)

var Val = uint16(math.MaxUint16 / 2)
var slice = make([]uint16, 500*500)

func BenchmarkRangeReadSliceByIndex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j, _ := range slice {
			Val = slice[j]
		}
	}
}

func BenchmarkRangeReadSliceByValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, v := range slice {
			Val = v
		}
	}
}

func BenchmarkRangeWriteSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j, _ := range slice {
			slice[j] = Val
		}
	}
}

func BenchmarkRangeReadAndWriteSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j, v := range slice {
			slice[j] = v
		}
	}
}

func BenchmarkForIterReadSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l := len(slice)
		for j := 0; j < l; j++ {
			Val = slice[j]
		}
	}
}

func BenchmarkForIterWriteSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l := len(slice)
		for j := 0; j < l; j++ {
			slice[j] = Val
		}
	}
}

func BenchmarkForIterReadAndWriteSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l := len(slice)
		for j := 0; j < l; j++ {
			slice[j] = slice[j]
		}
	}
}
