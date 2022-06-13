package gage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvenSucces(t *testing.T) {
	result := CheckNum(1, 2, 3, 4, 5)
	assert.Equal(t, "ganjil,genap,ganjil,genap,ganjil", result)
}
func BenchmarkTestEvenSuccess(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CheckNum(1, 2, 3, 4, 5)
	}
}

func TestEvenOddWithTable(t *testing.T) {
	test := []struct {
		name     string
		request  []int
		expected string
	}{
		{
			name:     "Check even number",
			request:  []int{0, 2, 4},
			expected: "genap, genap, genap",
		}, {
			name:     "Check odd number",
			request:  []int{1, 3, 5},
			expected: "ganjil,ganjil,ganjil",
		}, {
			name:     "Check even or odd",
			request:  []int{1, 2, 3, 4, 5},
			expected: "ganjil,genap,ganjil,genap,ganjil",
		},
	}

	for _, tes := range test {
		t.Run(tes.name, func(t *testing.T) {
			result := CheckNum(tes.request...)
			assert.Equal(t, tes.expected, result)
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
			expected: "genap,genap,genap",
		}, {
			name:     "Check odd number",
			request:  []int{1, 3, 5},
			expected: "ganjil,ganjil,ganjil",
		}, {
			name:     "Check even or odd",
			request:  []int{1, 2, 3, 4, 5},
			expected: "ganjil,genap,ganjil,genap,ganjil",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			CheckNum(benchmark.request...)
		})
	}
}
