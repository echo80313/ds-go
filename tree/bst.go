package tree

import (
	"fmt"
	"io"
)

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
			ptr = ptr.chld[RightChld]
		} else {
			ptr = ptr.chld[LeftChld]
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
			nxt = ptr.chld[RightChld]
			if nxt == nullBinarySearchTreeNode {
				newNode := NewBinarySearchTreeNode(data)
				newNode.parent = ptr
				newNode.childWhich = 1
				ptr.chld[RightChld] = newNode
				return
			}
		} else {
			nxt = ptr.chld[LeftChld]
			if nxt == nullBinarySearchTreeNode {
				newNode := NewBinarySearchTreeNode(data)
				newNode.parent = ptr
				newNode.childWhich = 0
				ptr.chld[LeftChld] = newNode
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
	if node.chld[LeftChld] == nullBinarySearchTreeNode &&
		node.chld[RightChld] == nullBinarySearchTreeNode {
		// leaf case
		node.parent.chld[node.childWhich] = nullBinarySearchTreeNode
		return nil
	}
	if node.chld[LeftChld] != nullBinarySearchTreeNode &&
		node.chld[RightChld] != nullBinarySearchTreeNode {
		// have both
		successor := b.inorderSuccessorHelper(node)
		node.data = successor.data
		successor.parent.chld[successor.childWhich] = nullBinarySearchTreeNode
		return nil
	}

	// has only one child
	for i := 0; i < 2; i++ {
		if node.chld[i] != nullBinarySearchTreeNode {
			node.parent.chld[node.childWhich] = node.chld[i]
			node.chld[i].parent = node.parent
			node.chld[i].childWhich = node.childWhich
			return nil
		}
	}
	return nil
}

func (b *BinarySearchTree) InOrderPrint(writer io.Writer) {
	stk := make([]*BinarySearchTreeNode, 0)
	cur := b.root
	for cur != nullBinarySearchTreeNode || len(stk) > 0 {
		for cur != nullBinarySearchTreeNode {
			stk = append(stk, cur)
			cur = cur.chld[LeftChld]
		}

		cur, stk = stk[len(stk)-1], stk[:len(stk)-1]
		fmt.Fprintf(writer, "%v ", cur.data)

		cur = cur.chld[RightChld]
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
	if node.chld[RightChld] != nullBinarySearchTreeNode {
		return b.findMost(node.chld[RightChld], 0)
	}
	cur := b.root
	for cur != nullBinarySearchTreeNode {
		if node.data.Less(cur.data) {
			return cur
		} else if !node.data.Equal(cur.data) {
			cur = cur.chld[RightChld]
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
