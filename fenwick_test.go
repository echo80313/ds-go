package ds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFenwickTreeSumAndUpdate(t *testing.T) {
	ft, err := NewFenwickTree(16)
	assert.Nil(t, err)
	assert.Equal(t, 0, ft.Sum(16))

	ft.Update(16, 10)
	assert.Equal(t, 0, ft.Sum(8))
	assert.Equal(t, 10, ft.Sum(16))

	ft.Update(8, 10)
	assert.Equal(t, 10, ft.Sum(8))
	assert.Equal(t, 20, ft.Sum(16))
}
