package union_find

import "errors"

var errInvalidIndex = errors.New("invalid index")

type UnionFind struct {
	n      int
	parent []int
	size   []int
}

func NewUnionFind(n int) (*UnionFind, error) {
	if n <= 0 {
		return nil, errors.New("size needs to be positive")
	}

	pa := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		pa[i] = i
		size[i] = 1
	}

	return &UnionFind{
		n:      n,
		parent: pa,
		size:   size,
	}, nil
}

func (u *UnionFind) isValidIndex(x int) bool {
	return x >= 0 && x < u.n
}

func (u *UnionFind) unionBySize(x, y int) {
	xRoot := u.findWithPathCompression(x)
	yRoot := u.findWithPathCompression(y)

	// Already in the same connected component, done
	if xRoot == yRoot {
		return
	}

	// attach small size subtree to larger one
	if u.size[xRoot] < u.size[yRoot] {
		xRoot, yRoot = yRoot, xRoot
	}
	u.parent[yRoot] = xRoot
	u.size[xRoot] += u.size[yRoot]
}

func (u *UnionFind) findWithPathCompression(x int) int {
	if u.parent[x] == x {
		return x
	}
	u.parent[x] = u.findWithPathCompression(u.parent[x])
	return u.parent[x]
}

/*
External API: Union and Find
*/

func (u *UnionFind) Find(x int) (int, error) {
	if !u.isValidIndex(x) {
		return -1, errInvalidIndex
	}
	return u.findWithPathCompression(x), nil
}

func (u *UnionFind) Union(x, y int) error {
	if !u.isValidIndex(x) || !u.isValidIndex(y) {
		return errInvalidIndex
	}
	u.unionBySize(x, y)
	return nil
}
