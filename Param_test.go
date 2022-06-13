package gage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvenSucces(t *testing.T) {
	result := CheckNum(3)
	assert.Equal(t, "genap", result, "Test result must be 'genap' ")
}

func TestOddFailed(t *testing.T) {
	result := CheckNum(2)
	assert.Equal(t, "ganjil", result, "Test result must be 'ganjil' ")
}

func TestEvenOddWithTable(t *testing.T) {
	test := []struct {
		name     string
		request  int
		expected string
	}{
		{
			name:     "Check even number",
			request:  2,
			expected: "genap",
		}, {
			name:     "Check odd number",
			request:  1,
			expected: "ganjil",
		}, {
			name:     "Check zero",
			request:  0,
			expected: "genap",
		},
	}

	for _, tes := range test {
		t.Run(tes.name, func(t *testing.T) {
			result := CheckNum(tes.request)
			assert.Equal(t, tes.expected, result)
		})
	}
}
