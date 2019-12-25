package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var minHeap = NewMinHeap()

var nodes = [8]int{10, 9, 8, 7, 6, 3, 5, 4}

func TestMinHeapInsert(t *testing.T) {
	require.Equal(t, -1, minHeap.extractMin())
	for _, n := range nodes {
		minHeap.insert(n)
	}
	require.Equal(t, 3, minHeap.extractMin())
}
