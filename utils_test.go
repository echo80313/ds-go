package ds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLSB(t *testing.T) {
	assert.Equal(t, 0, LeastSignificantBit(0))
	assert.Equal(t, 1, LeastSignificantBit(1))
	assert.Equal(t, 2, LeastSignificantBit(22))
}

func TestNextPowerOf2(t *testing.T) {
	assert.Equal(t, 1, NextPowerOf2(0))
	assert.Equal(t, 1, NextPowerOf2(1))
	assert.Equal(t, 2, NextPowerOf2(2))
	assert.Equal(t, 32, NextPowerOf2(17))
	assert.Equal(t, 128, NextPowerOf2(127))
	assert.Equal(t, 128, NextPowerOf2(128))
	assert.Equal(t, 1<<31, NextPowerOf2((1<<31)-1))
	assert.Equal(t, 1<<31, NextPowerOf2(1<<31))
}
