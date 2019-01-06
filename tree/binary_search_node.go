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

const (
	LeftChld  = 0
	RightChld = 1
)

// BinarySearchTreeNode is type that used by all
// kinds of binary search tree.
type BinarySearchTreeNode struct {
	data       ComparableValue
	chld       []*BinarySearchTreeNode
	parent     *BinarySearchTreeNode
	childWhich int
}

func NewBinarySearchTreeNode(data ComparableValue) *BinarySearchTreeNode {
	node := &BinarySearchTreeNode{
		data: data,
		chld: make([]*BinarySearchTreeNode, 2),
	}
	node.chld[LeftChld] = nullBinarySearchTreeNode
	node.chld[RightChld] = nullBinarySearchTreeNode
	node.parent = nullBinarySearchTreeNode
	return node
}

// Rotate performes, yes, bst rotation.
// d = 0 Left rotate
// d = 1 Right rotate
func (o *BinarySearchTreeNode) Rotate(d int) {
	k := o.chld[d^1]
	oParent := o.parent
	oChldWhich := o.childWhich

	o.chld[d^1] = k.chld[d]
	k.chld[d].parent = o.chld[d^1]
	k.chld[d].childWhich = d

	k.chld[d] = o
	o.parent = k.chld[d]
	o.childWhich = d

	oParent.chld[oChldWhich] = k
}

var nullBinarySearchTreeNode = &BinarySearchTreeNode{}
