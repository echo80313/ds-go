package tree

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
	data ComparableValue
	chld []*BinarySearchTreeNode
}

var nullBinarySearchTreeNode = &BinarySearchTreeNode{}

func NewBinarySearchTreeNode(data ComparableValue) *BinarySearchTreeNode {
	node := &BinarySearchTreeNode{
		data: data,
		chld: make([]*BinarySearchTreeNode, 2),
	}
	node.chld[0] = nullBinarySearchTreeNode
	node.chld[1] = nullBinarySearchTreeNode
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
	ptr := b.root
	for ptr != nullBinarySearchTreeNode {
		if ptr.data.Equal(data) {
			return true
		}
		if ptr.data.Less(data) {
			ptr = ptr.chld[0]
		} else {
			ptr = ptr.chld[1]
		}
	}
	return false
}

func (b *BinarySearchTree) put(data ComparableValue, ptr *BinarySearchTreeNode, dir int) {
	ptr.chld[dir] = NewBinarySearchTreeNode(data)
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
			nxt = ptr.chld[0]
			if nxt == nullBinarySearchTreeNode {
				ptr.chld[0] = NewBinarySearchTreeNode(data)
				return
			}
		} else {
			nxt = ptr.chld[1]
			if nxt == nullBinarySearchTreeNode {
				ptr.chld[1] = NewBinarySearchTreeNode(data)
				return
			}
		}
		ptr = nxt
	}
}
