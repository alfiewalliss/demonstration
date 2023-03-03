package factorial

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnitFactorial(t *testing.T) {
	assert := assert.New(t)
	tests :=
		[]struct {
			n      int
			result int
		}{
			{0, 1},
			{1, 1},
			{2, 2},
			{5, 10},
			{6, 720},
		}

	for _, test := range tests {
		/* act */
		v := Factorial(tesxt.n)
		/* assert */
		assert.Equal(test.result, v)
	}
}
