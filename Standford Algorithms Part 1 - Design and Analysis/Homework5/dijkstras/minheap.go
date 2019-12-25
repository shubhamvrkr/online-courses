package main

import (
	"errors"
	"fmt"
)

var MAX_INFINITY int = 999999999999999
var MIN_INFINITY int = -99999999999999

type Node struct {
	label string
	key   int
}

type MinHeap struct {
	size int
	data []Node
	pos  map[string]int
}

func NewMinHeap() *MinHeap {
	return &MinHeap{-1, []Node{}, make(map[string]int)}
}

func (h *MinHeap) insert(key Node) {

	if h.isEmpty() {
		h.size++
		list := []Node{key}
		h.data = list
		h.pos[key.label] = h.size
	} else {
		h.size++
		h.data = append(h.data, key)
		h.pos[key.label] = h.size
	}
	h.heapify(h.size)
}

//restrict calling heapify from outside method
func (h *MinHeap) heapify(index int) {
	p := getParentIndex(index)
	for index > 0 && h.data[index].key < h.data[p].key {
		h.data[index], h.data[p] = h.data[p], h.data[index]
		h.pos[h.data[index].label] = index
		h.pos[h.data[p].label] = p

		index = p
		p = getParentIndex(index)
	}
}

func (h *MinHeap) extractMin() *Node {
	if h.isEmpty() {
		return nil
	}
	min := h.data[0]
	h.pos[min.label] = -1
	h.data[0] = h.data[h.size]
	h.pos[h.data[0].label] = 0
	h.data = h.data[:h.size]
	h.size--
	parent := 0
	smallestIndex, err := h.findSmallestChildren(parent)
	for err == nil {
		if h.data[smallestIndex].key < h.data[parent].key {
			h.data[parent], h.data[smallestIndex] = h.data[smallestIndex], h.data[parent]
			h.pos[h.data[parent].label] = parent
			h.pos[h.data[smallestIndex].label] = smallestIndex
			parent = smallestIndex
			smallestIndex, err = h.findSmallestChildren(parent)
		} else {
			break
		}
	}
	return &min
}

//incase of duplicates first value found will be replaced with the new value
func (h *MinHeap) decreaseKey(label string, key int) {
	k := h.pos[label]
	node := h.data[k]
	node.key = key
	h.data[k] = node
	h.heapify(k)
}

func (h *MinHeap) print() {
	for _, val := range h.data {
		fmt.Println(val, " ")
	}
}

func (h *MinHeap) getSize() int {
	return h.size + 1
}

func (h *MinHeap) getNodeByLabel(label string) *Node {
	pos := h.pos[label]
	if pos == -1 {
		return nil
	}
	return &h.data[pos]
}
func (h *MinHeap) isEmpty() bool {
	if h.size == -1 {
		return true
	}
	return false
}

func getParentIndex(index int) int {
	if index%2 == 0 {
		return index/2 - 1
	} else {
		return index / 2
	}
}

func getLeftChildrenIndex(index int) int {
	return 2*index + 1
}

func getRightChildrenIndex(index int) int {
	return 2*index + 2
}

func (h *MinHeap) findSmallestChildren(p int) (int, error) {
	leftChildIndex := getLeftChildrenIndex(p)
	rightChildIndex := getRightChildrenIndex(p)
	if leftChildIndex <= h.size && rightChildIndex <= h.size {
		if h.data[leftChildIndex].key < h.data[rightChildIndex].key {
			return leftChildIndex, nil
		} else {
			return rightChildIndex, nil
		}
	} else if leftChildIndex <= h.size && rightChildIndex > h.size {
		//doesnt contains right child
		return leftChildIndex, nil
	} else if leftChildIndex > h.size && rightChildIndex <= h.size {
		//doesnt contains left child
		return rightChildIndex, nil
	} else {
		return -1, errors.New("No children present for the node")
	}
}
