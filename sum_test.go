package algorithms_in_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSum(t *testing.T) {
	x := Sum([]int{1,2,3})
	assert.True(t, x==6)
}
