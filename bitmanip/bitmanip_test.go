package bitmanip

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Base2(t *testing.T) {
	testCases := []struct {
		in  string
		out uint64
	}{
		{"0", 0},
		{"1", 1},
		{"00001", 1},
		{"00010", 2},
		{"00101100101011000", 22872},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc.in), func(t *testing.T) {
			res := Base2(tc.in)
			assert.Equal(t, tc.out, res)
		})
	}
}

var testCases = []struct {
		in  string
		out uint64
	}{
		{"00100", 1},
		{"00000", 0},
		{"00111", 1},
		{"00101", 0},
		{"001000000000001", 0},
		{"0010000011100001", 1},
		{"011101010100101010111010001010101011000001010101010101010101111", 0},
		{"0111010101001010101110100010101010110000010101010101010101011110", 0},
	}
var shared_in uint64 = Base2("011101010100101010111010001010101011000001010101010101010101111")


func Test_Parity(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc.in), func(t *testing.T) {
			res := Parity(Base2(tc.in))
			assert.Equal(t, tc.out, res)
		})
	}
}


func Benchmark_Parity(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Parity(shared_in)
	}
}

func Test_FasterParity(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc.in), func(t *testing.T) {
			res := FasterParity(Base2(tc.in))
			assert.Equal(t, tc.out, res)
		})
	}
}

func Benchmark_FasterParity(b *testing.B) {
	for n := 0; n < b.N; n++ {
		FasterParity(shared_in)
	}
}

func Test_EvenFasterParity(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc.in), func(t *testing.T) {
			res := EvenFasterParity(Base2(tc.in))
			assert.Equal(t, tc.out, res)
		})
	}
}

func Benchmark_EvenFasterParity(b *testing.B) {
	for n := 0; n < b.N; n++ {
		EvenFasterParity(shared_in)
	}
}

func Test_CachedParity(t *testing.T) {
	cache := ProduceCache()

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc.in), func(t *testing.T) {
			res := CachedParity(Base2(tc.in), cache)
			assert.Equal(t, tc.out, res)
		})
	}
}

func Benchmark_CachedParity(b *testing.B) {
	cache := ProduceCache()
	for n := 0; n < b.N; n++ {
		CachedParity(shared_in, cache)
	}
}

func Test_XORParity(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc.in), func(t *testing.T) {
			res := XORParity(Base2(tc.in))
			assert.Equal(t, tc.out, res)
		})
	}
}

func Benchmark_XORParity(b *testing.B) {
	for n := 0; n < b.N; n++ {
		XORParity(shared_in)
	}
}