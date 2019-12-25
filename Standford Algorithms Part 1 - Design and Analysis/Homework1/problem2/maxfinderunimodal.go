// Problem Statement:
// You are a given a unimodal array of n distinct elements,
// meaning that its entries are in increasing order up until its maximum element, after which its elements are in decreasing order.
// Give an algorithm to compute the maximum element that runs in O(log n) time.

package main

import "fmt"

func main() {

	//Given Facts :
	//N distinct elements
	var input1 = []int{1, 2, 3, 4, 20, 19, 16, 15, 13, 12, 11, 10, 9, 8}
	var input2 = []int{15, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	var input3 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	largest := getMaximun(input1, 0, len(input1)-1)
	fmt.Println("Largest: ", largest)
	largest = getMaximun(input2, 0, len(input2)-1)
	fmt.Println("Largest: ", largest)
	largest = getMaximun(input3, 0, len(input3)-1)
	fmt.Println("Largest: ", largest)

}

func getMaximun(input []int, a int, b int) int {
	if a == b {
		return input[a]
	} else if a+1 == b {
		if input[a] > input[b] {
			return input[a]
		}
		return input[b]
	} else {
		mid := (a + b) / 2
		if input[mid-1] < input[mid] {
			return getMaximun(input, mid, b)
		} else {
			return getMaximun(input, a, mid)
		}
	}
}
