package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var minHeap = NewMinHeap()

var nodes = [8]Node{Node{"0", 10}, Node{"1", 9}, Node{"2", 8}, Node{"3", 7}, Node{"4", 6},
	Node{"5", 5}, Node{"6", 4}, Node{"7", 3}}

func TestMinHeapInsert(t *testing.T) {
	var empty *Node
	require.Equal(t, empty, minHeap.extractMin())
	for _, n := range nodes {
		minHeap.insert(n)
	}
	require.Equal(t, "7", minHeap.extractMin().label)
}

func TestMinHeapDecreaseKey(t *testing.T) {
	size := minHeap.getSize()
	minHeap.decreaseKey("0", 1)
	minHeap.print()
	require.Equal(t, "0", minHeap.extractMin().label)
	require.Equal(t, size-1, minHeap.getSize())
	require.Equal(t, 4, minHeap.getNodeByLabel("6").key)
	require.Equal(t, 5, minHeap.getNodeByLabel("5").key)
	require.Equal(t, 9, minHeap.getNodeByLabel("1").key)

}
