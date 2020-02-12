package main

import "errors"

type UnionFind struct {
	ids     []int
	sz      []int
	leaders int
}

var initNumberNotPositive = errors.New("number of elements should be positive")

var indexOutOfRange = errors.New("index out of range")

// Create a new UnionFind structure with size n.
func New(n int) (*UnionFind, error) {
	if n < 0 {
		return nil, initNumberNotPositive
	}
	ids := make([]int, n+1)
	sz := make([]int, n+1)
	// Set id of each item to itself
	for i := 1; i <= n; i++ {
		ids[i] = i
	}
	return &UnionFind{ids: ids, sz: sz, leaders: n}, nil
}

// Interpretation of id[i]: id[i] is parent of i
// Root of i is id[id[id[...id[i]...]]]
func (uf *UnionFind) Root(i int) (int, error) {
	err := uf.checkIndexRange(i)
	if err != nil {
		return 0, err
	}

	// Chase parent pointers until reach root
	for i != uf.ids[i] {
		// Path compression
		// Make every other node in path point to its grandparent (thereby halving path length)
		uf.ids[i] = uf.ids[uf.ids[i]]
		i = uf.ids[i]
	}
	return i, nil
}

// Add connection between p and q
func (uf *UnionFind) Union(p int, q int) error {
	err := uf.checkIndexRange(p, q)
	if err != nil {
		return err
	}

	i, _ := uf.Root(p)
	j, _ := uf.Root(q)
	if i == j {
		return nil
	}
	if uf.sz[i] < uf.sz[j] {
		uf.ids[i] = j
		uf.sz[j] += uf.sz[i]
	} else {
		uf.ids[j] = i
		uf.sz[i] += uf.sz[j]
	}
	uf.leaders = uf.leaders - 1
	return nil
}

// Check whether p and q are in the same component
func (uf *UnionFind) Connected(p int, q int) (bool, error) {
	err := uf.checkIndexRange(p, q)
	if err != nil {
		return false, err
	}
	pRoot, _ := uf.Root(p)
	qRoot, _ := uf.Root(q)
	return pRoot == qRoot, nil
}

func (uf *UnionFind) checkIndexRange(indexes ...int) error {
	for _, i := range indexes {
		if i >= len(uf.ids) || i < 0 {
			return indexOutOfRange
		}
	}
	return nil
}
