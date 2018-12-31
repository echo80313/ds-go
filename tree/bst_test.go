package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBSTInsertAndDelete(t *testing.T) {
	bst := NewBinarySearchTree()
	bst.Insert(Int(0))
	assert.True(t, bst.Find(Int(0)))
	bst.Insert(Int(1))
	assert.True(t, bst.Find(Int(1)))
	assert.False(t, bst.Find(Int(5)))
}

type testWriter struct {
	buffer []byte
}

func (tw *testWriter) Write(p []byte) (int, error) {
	tw.buffer = append(tw.buffer, p...)
	return len(p), nil
}

func TestBSTInorderTraversal(t *testing.T) {
	bst := NewBinarySearchTree()
	bst.Insert(Int(0))
	bst.Insert(Int(10))
	bst.Insert(Int(8))
	bst.Insert(Int(12))
	bst.Insert(Int(1))
	bst.Insert(Int(19))

	tw := &testWriter{buffer: make([]byte, 0)}
	bst.InOrderPrint(tw)
	assert.Equal(t, []byte("0 1 8 10 12 19 \n"), tw.buffer)
}
