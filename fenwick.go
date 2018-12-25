package ds

import "fmt"

// FenwickTree
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
		data: make([]int, NextPowerOf2(sz+1)),
	}, nil
}

func (f *FenwickTree) Update(at, val int) {
	for at <= f.size {
		f.data[at] += val
		at += LeastSignificantBit(at)
	}
}

func (f *FenwickTree) Sum(at int) int {
	sum := 0
	for at > 0 {
		sum += f.data[at]
		at -= LeastSignificantBit(at)
	}
	return sum
}
