package main

import (
	"errors"
	"fmt"
)

// var MAX_INFINITY int = 999999999999999
// var MIN_INFINITY int = -99999999999999

type MaxHeap struct {
	size int
	data []int
}

func (h *MaxHeap) viewMax() int {
	if h.isEmpty() {
		return -1
	}
	return h.data[0]
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{-1, []int{}}
}

func (h *MaxHeap) insert(key int) {
	if h.isEmpty() {
		h.size++
		list := []int{key}
		h.data = list
	} else {
		h.size++
		h.data = append(h.data, key)
	}
	h.heapify(h.size)
}

//restrict calling heapify from outside method
func (h *MaxHeap) heapify(index int) {
	p := getParentIndex(index)
	for index > 0 && h.data[index] > h.data[p] {
		h.data[index], h.data[p] = h.data[p], h.data[index]
		index = p
		p = getParentIndex(index)
	}
}

func (h *MaxHeap) extractMax() int {
	if h.isEmpty() {
		return -1
	}
	min := h.data[0]
	h.data[0] = h.data[h.size]
	h.data = h.data[:h.size]
	h.size--
	parent := 0
	smallestIndex, err := h.findSmallestChildren(parent)
	for err == nil {
		if h.data[smallestIndex] > h.data[parent] {
			h.data[parent], h.data[smallestIndex] = h.data[smallestIndex], h.data[parent]
			parent = smallestIndex
			smallestIndex, err = h.findSmallestChildren(parent)
		} else {
			break
		}
	}
	return min
}

func (h *MaxHeap) print() {
	fmt.Print("MaxHeap: ")
	for _, val := range h.data {
		fmt.Print(val, " ")
	}
	fmt.Println()
}

func (h *MaxHeap) getSize() int {
	return h.size + 1
}

func (h *MaxHeap) isEmpty() bool {
	if h.size == -1 {
		return true
	}
	return false
}

func (h *MaxHeap) findSmallestChildren(p int) (int, error) {
	leftChildIndex := getLeftChildrenIndex(p)
	rightChildIndex := getRightChildrenIndex(p)
	if leftChildIndex <= h.size && rightChildIndex <= h.size {
		if h.data[leftChildIndex] > h.data[rightChildIndex] {
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

// func getParentIndex(index int) int {
// 	if index%2 == 0 {
// 		return index/2 - 1
// 	} else {
// 		return index / 2
// 	}
// }
//
// func getLeftChildrenIndex(index int) int {
// 	return 2*index + 1
// }
//
// func getRightChildrenIndex(index int) int {
// 	return 2*index + 2
// }
