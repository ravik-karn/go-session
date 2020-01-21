package table

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	tests := []struct {
		name string
		num1 int
		num2 int
		res  int
	}{
		{"1-plus-1-is-2", 1, 1, 2},
		{"2-plus-2-is-4", 2, 2, 4},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.res, Sum(test.num1, test.num2))
		})
	}
}

func TestSum1(t *testing.T) {
	res := Sum(1, 2)
	assert.Equal(t, 3, res)
}
