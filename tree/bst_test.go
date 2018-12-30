package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBST(t *testing.T) {
	bst := NewBinarySearchTree()
	bst.Insert(Int(0))
	assert.True(t, bst.Find(Int(0)))
	bst.Insert(Int(1))
	assert.True(t, bst.Find(Int(1)))
	assert.False(t, bst.Find(Int(5)))
}
