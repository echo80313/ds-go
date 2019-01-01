package tree

import (
	"fmt"
	"io"
)

type ComparableValue interface {
	Less(interface{}) bool
	Equal(interface{}) bool
}

type Int int

func (i Int) Less(v interface{}) bool {
	if vInt, ok := v.(Int); ok {
		return i < vInt
	}
	return false
}

func (i Int) Equal(v interface{}) bool {
	if vInt, ok := v.(Int); ok {
		return i == vInt
	}
	return false
}

var _ ComparableValue = Int(0)

type BinarySearchTreeNode struct {
	data       ComparableValue
	chld       []*BinarySearchTreeNode
	parent     *BinarySearchTreeNode
	childWhich int
}

var nullBinarySearchTreeNode = &BinarySearchTreeNode{}

func NewBinarySearchTreeNode(data ComparableValue) *BinarySearchTreeNode {
	node := &BinarySearchTreeNode{
		data: data,
		chld: make([]*BinarySearchTreeNode, 2),
	}
	node.chld[0] = nullBinarySearchTreeNode
	node.chld[1] = nullBinarySearchTreeNode
	node.parent = nullBinarySearchTreeNode
	return node
}

// BinarySearchTree yes, bst
type BinarySearchTree struct {
	root *BinarySearchTreeNode
}

func NewBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{root: nullBinarySearchTreeNode}
}

func (b *BinarySearchTree) Find(data ComparableValue) bool {
	return b.findNode(data) != nullBinarySearchTreeNode
}

func (b *BinarySearchTree) findNode(data ComparableValue) *BinarySearchTreeNode {
	ptr := b.root
	for ptr != nullBinarySearchTreeNode {
		if ptr.data.Equal(data) {
			break
		}
		if ptr.data.Less(data) {
			ptr = ptr.chld[1]
		} else {
			ptr = ptr.chld[0]
		}
	}
	return ptr
}

func (b *BinarySearchTree) Insert(data ComparableValue) {
	if b.root == nullBinarySearchTreeNode {
		b.root = NewBinarySearchTreeNode(data)
		return
	}
	ptr := b.root
	var nxt *BinarySearchTreeNode
	for {
		if ptr.data.Less(data) {
			nxt = ptr.chld[1]
			if nxt == nullBinarySearchTreeNode {
				newNode := NewBinarySearchTreeNode(data)
				newNode.parent = ptr
				newNode.childWhich = 1
				ptr.chld[1] = newNode
				return
			}
		} else {
			nxt = ptr.chld[0]
			if nxt == nullBinarySearchTreeNode {
				newNode := NewBinarySearchTreeNode(data)
				newNode.parent = ptr
				newNode.childWhich = 0
				ptr.chld[0] = newNode
				return
			}
		}
		ptr = nxt
	}
}

func (b *BinarySearchTree) Delete(data ComparableValue) error {
	node := b.findNode(data)
	if node == nullBinarySearchTreeNode {
		return fmt.Errorf("Value: %v doesn't exist", data)
	}
	if node.chld[0] == nullBinarySearchTreeNode &&
		node.chld[1] == nullBinarySearchTreeNode {
		// leaf case
		node.parent.chld[node.childWhich] = nullBinarySearchTreeNode
		return nil
	}
	if node.chld[0] != nullBinarySearchTreeNode &&
		node.chld[1] != nullBinarySearchTreeNode {
		// have both
		successor := b.inorderSuccessorHelper(node)
		node.data = successor.data
		successor.parent.chld[successor.childWhich] = nullBinarySearchTreeNode
		return nil
	}

	// has only one child
	if node.chld[0] != nullBinarySearchTreeNode {
		node.parent.chld[node.childWhich] = node.chld[0]
		node.chld[0].parent = node.parent
		node.chld[0].childWhich = node.childWhich
		return nil
	}
	node.parent.chld[node.childWhich] = node.chld[1]
	node.chld[1].parent = node.parent
	node.chld[1].childWhich = node.childWhich
	return nil
}

func (b *BinarySearchTree) InOrderPrint(writer io.Writer) {
	stk := make([]*BinarySearchTreeNode, 0)
	cur := b.root
	for cur != nullBinarySearchTreeNode || len(stk) > 0 {
		for cur != nullBinarySearchTreeNode {
			stk = append(stk, cur)
			cur = cur.chld[0]
		}

		cur, stk = stk[len(stk)-1], stk[:len(stk)-1]
		fmt.Fprintf(writer, "%v ", cur.data)

		cur = cur.chld[1]
	}
	fmt.Fprintf(writer, "\n")
}

func (b *BinarySearchTree) findMost(node *BinarySearchTreeNode, polar int) *BinarySearchTreeNode {
	cur := node
	for cur.chld[polar] != nullBinarySearchTreeNode {
		cur = cur.chld[polar]
	}
	return cur
}

func (b *BinarySearchTree) FindMax() ComparableValue {
	return b.findMost(b.root, 1).data
}

func (b *BinarySearchTree) FindMin() ComparableValue {
	return b.findMost(b.root, 0).data
}

func (b *BinarySearchTree) inorderSuccessorHelper(
	node *BinarySearchTreeNode) *BinarySearchTreeNode {
	if node.chld[1] != nullBinarySearchTreeNode {
		return b.findMost(node.chld[1], 0)
	}
	cur := b.root
	for cur != nullBinarySearchTreeNode {
		if node.data.Less(cur.data) {
			return cur
		} else if !node.data.Equal(cur.data) {
			cur = cur.chld[1]
		} else {
			// sad...
			return nullBinarySearchTreeNode
		}
	}
	// should not reach here bcz node is not in the tree
	// in this case
	return nullBinarySearchTreeNode
}

func (b *BinarySearchTree) InorderSuccessor(data ComparableValue) (ComparableValue, error) {
	node := b.findNode(data)
	if node == nullBinarySearchTreeNode {
		return nil, fmt.Errorf("InorderSuccessor: can't find node with value %v", data)
	}
	successor := b.inorderSuccessorHelper(node)
	if successor != nullBinarySearchTreeNode {
		return successor.data, nil
	}
	return nil, fmt.Errorf("InorderSuccessor: %v is the last value", data)
}
