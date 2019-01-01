package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBSTInsertAndFind(t *testing.T) {
	bst := NewBinarySearchTree()
	bst.Insert(Int(0))
	assert.True(t, bst.Find(Int(0)))
	bst.Insert(Int(1))
	assert.True(t, bst.Find(Int(1)))
	assert.False(t, bst.Find(Int(5)))

	for i := 0; i < 100; i++ {
		bst.Insert(Int(i))
	}
	for i := 0; i < 100; i++ {
		assert.True(t, bst.Find(Int(i)))
	}
}

func TestFindMinAndMax(t *testing.T) {
	bst := NewBinarySearchTree()
	bst.Insert(Int(0))
	bst.Insert(Int(10))
	bst.Insert(Int(8))
	bst.Insert(Int(12))
	bst.Insert(Int(1))
	bst.Insert(Int(19))

	assert.Equal(t, Int(19), bst.FindMax())
	assert.Equal(t, Int(0), bst.FindMin())
}

func TestFindInorderSuccessor(t *testing.T) {
	bst := NewBinarySearchTree()
	bst.Insert(Int(10))
	bst.Insert(Int(8))
	bst.Insert(Int(12))
	bst.Insert(Int(1))
	bst.Insert(Int(19))
	bst.Insert(Int(0))

	s, err := bst.InorderSuccessor(Int(10))
	assert.Nil(t, err)
	assert.Equal(t, Int(12), s)

	s, err = bst.InorderSuccessor(Int(8))
	assert.Nil(t, err)
	assert.Equal(t, Int(10), s)

	s, err = bst.InorderSuccessor(Int(19))
	assert.NotNil(t, err)
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
