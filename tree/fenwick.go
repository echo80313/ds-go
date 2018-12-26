package tree

import (
	"fmt"

	ds "github.com/echo80313/ds-go"
)

// FenwickTree a.k.a BIT tree supports point add and prefix sum
// query in O(n log n) which n is the array size embedded in FenwickTree
// It's 1-indexed
type FenwickTree struct {
	size int
	data []int
}

func NewFenwickTree(sz int) (*FenwickTree, error) {
	if sz <= 0 {
		return nil, fmt.Errorf("Invalid size: %d", sz)
	}
	return &FenwickTree{
		size: sz,
		data: make([]int, ds.NextPowerOf2(sz+1)),
	}, nil
}

// Add adds value to element at index `at`
func (f *FenwickTree) Add(at, val int) {
	for at <= f.size {
		f.data[at] += val
		at += ds.LeastSignificantBit(at)
	}
}

// Sum calculates the prefix sum [1 ... at]
func (f *FenwickTree) Sum(at int) int {
	sum := 0
	for at > 0 {
		sum += f.data[at]
		at -= ds.LeastSignificantBit(at)
	}
	return sum
}
