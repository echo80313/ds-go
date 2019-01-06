package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotate(t *testing.T) {
	bst := NewBinarySearchTree()
	bst.Insert(Int(2))
	bst.Insert(Int(1))
	bst.Insert(Int(3))
	bst.root.chld[LeftChld].Rotate(LeftChld) // Left rotate
	assert.Equal(t, Int(3), bst.root.chld[LeftChld].data)
	assert.Equal(t, Int(2), bst.root.chld[LeftChld].chld[LeftChld].data)
	assert.Equal(t, nullBinarySearchTreeNode, bst.root.chld[LeftChld].chld[RightChld])
	assert.Equal(t, Int(1), bst.root.chld[LeftChld].chld[LeftChld].chld[LeftChld].data)
	assert.Equal(t, nullBinarySearchTreeNode, bst.root.chld[LeftChld].chld[LeftChld].chld[RightChld])

	bst.root.chld[LeftChld].Rotate(RightChld)
	assert.Equal(t, Int(2), bst.root.chld[LeftChld].data)
	assert.Equal(t, Int(1), bst.root.chld[LeftChld].chld[LeftChld].data)
	assert.Equal(t, Int(3), bst.root.chld[LeftChld].chld[RightChld].data)
}
