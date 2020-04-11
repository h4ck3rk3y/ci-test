package foo

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestSum(t *testing.T) {
	total := Sum(5, 3)
	assert.Equal(t, total, 8)
}
