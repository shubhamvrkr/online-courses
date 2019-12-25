package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var maxHeap = NewMaxHeap()

var maxNodes = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 20, 19, 18, 17, 16, 15}

func TestMaxHeapInsert(t *testing.T) {
	require.Equal(t, -1, maxHeap.extractMax())
	for _, n := range maxNodes {
		maxHeap.insert(n)
	}
	require.Equal(t, 20, maxHeap.extractMax())
}
