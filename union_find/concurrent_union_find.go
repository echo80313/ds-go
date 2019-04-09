package union_find

import (
	"sync"
)

type ConcurrentUnionFind struct {
	sync.Mutex
	unionFind *UnionFind
}

func NewConcurrentUnionFind(n int) (*ConcurrentUnionFind, error) {
	uFind, err := NewUnionFind(n)
	if err != nil {
		return nil, err
	}
	return &ConcurrentUnionFind{
		unionFind: uFind,
	}, nil
}

func (cu *ConcurrentUnionFind) Union(x, y int) error {
	cu.Lock()
	defer cu.Unlock()
	return cu.unionFind.Union(x, y)
}

func (cu *ConcurrentUnionFind) Find(x int) (int, error) {
	cu.Lock()
	defer cu.Unlock()
	return cu.unionFind.Find(x)
}
