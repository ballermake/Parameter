package gage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvenSucces(t *testing.T) {
	result := CheckNum(1, 2, 3, 4, 5)
	assert.Equal(t, "GANJIL, GENAP, GANJIL, GENAP, GANJIL", result)
}
func BenchmarkTestEvenSuccess(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CheckNum(1, 2, 3, 4, 5)
	}
}

func TestEvenOddWithTable(t *testing.T) {
	tests := []struct {
		name     string
		request  []int
		expected string
	}{
		{
			name:     "Check even number",
			request:  []int{0, 2, 4},
			expected: "GENAP, GENAP, GENAP",
		}, {
			name:     "Check odd number",
			request:  []int{1, 3, 5},
			expected: "GANJIL, GANJIL, GANJIL",
		}, {
			name:     "Check even or odd",
			request:  []int{1, 2, 3, 4, 5},
			expected: "GANJIL, GENAP, GANJIL, GENAP, GANJIL",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := CheckNum(test.request...)
			assert.Equal(t, test.expected, result)
		})
	}
}

func BenchmarkTestEvenOddWithTable(b *testing.B) {
	benchmarks := []struct {
		name     string
		request  []int
		expected string
	}{
		{
			name:     "Check even number",
			request:  []int{0, 2, 4},
			expected: "GENAP, GENAP, GENAP",
		}, {
			name:     "Check odd number",
			request:  []int{1, 3, 5},
			expected: "GANJIL, GANJIL, GANJIL",
		}, {
			name:     "Check even or odd",
			request:  []int{1, 2, 3, 4, 5},
			expected: "GANJIL, GENAP, GANJIL, GENAP, GANJIL",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			CheckNum(benchmark.request...)
		})
	}
}
