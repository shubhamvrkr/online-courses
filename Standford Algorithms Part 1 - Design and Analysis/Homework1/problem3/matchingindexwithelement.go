// Problem Statement:
// You are given a sorted (from smallest to largest) array A of n distinct integers which can be positive, negative, or zero.
// You want to decide whether or not there is an index i such that A[i] = i.
// Design the fastest algorithm that you can for solving this problem.

package main

import "fmt"

func main() {

	// Given Facts :
	// Sorted array with +ve, -ve or zero
	// N Distinct elements
	var input1 = []int{-1, -2, -3, 3, 6, 7}
	var input2 = []int{1, 2, 3, 4, 5, 6}

	index := getMatchingIndex(input1, 0, len(input1)-1)
	fmt.Println("Index: ", index)
	index = getMatchingIndex(input2, 0, len(input2)-1)
	fmt.Println("Index: ", index)

}

func getMatchingIndex(input []int, a int, b int) int {
	if a == b && input[a] != a || b < a {
		return -1
	} else {
		mid := (a + b) / 2
		if input[mid] == mid {
			return mid
		} else if input[mid] < mid {
			return getMatchingIndex(input, mid+1, b)
		} else {
			return getMatchingIndex(input, a, mid-1)
		}
	}
}
