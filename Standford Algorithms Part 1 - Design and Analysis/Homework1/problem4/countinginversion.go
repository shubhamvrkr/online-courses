// Problem Statement:
// You are given a sorted (from smallest to largest) array A of n distinct integers which can be positive, negative, or zero.
// You want to decide whether or not there is an index i such that A[i] = i.
// Design the fastest algorithm that you can for solving this problem.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	// Given Facts:
	// distinct numbers
	var filename = "test.txt"
	var numbers = readFile(filename)
	invcount := countInversion(numbers, 0, len(numbers)-1)
	fmt.Println("Inversion Count: ", invcount)
}

func countInversion(input []int, a int, b int) int {

	if a == b {
		return 0
	} else if a+1 == b {
		if input[a] > input[b] {
			input[a], input[b] = input[b], input[a]
			return 1
		} else {
			return 0
		}
	} else {
		mid := (a + b) / 2
		leftInvCount := countInversion(input, a, mid)
		rightInvCount := countInversion(input, mid+1, b)
		splitCount := countSplitInversion(input, a, mid, b)
		return leftInvCount + splitCount + rightInvCount
	}
}

func countSplitInversion(input []int, a int, mid int, b int) int {
	var cacheA = a
	var cacheB = b
	var temp = []int{}
	count := 0
	j := mid + 1

	for a < mid+1 || j < b+1 {

		if a == mid+1 {
			temp = append(temp, input[j])
			j = j + 1

		} else if j == b+1 {
			temp = append(temp, input[a])
			a = a + 1
		} else {
			if input[a] > input[j] {
				count = count + mid + 1 - a
				temp = append(temp, input[j])
				j = j + 1
			} else {
				temp = append(temp, input[a])
				a = a + 1
			}
		}
	}
	for z, i := 0, cacheA; i <= cacheB; z, i = z+1, i+1 {
		input[i] = temp[z]
	}
	return count
}

func readFile(filename string) []int {
	var numbers = []int{}
	file, _ := os.Open(filename)
	defer file.Close()
	fscanner := bufio.NewScanner(file)
	for fscanner.Scan() {
		i, _ := strconv.Atoi(fscanner.Text())
		numbers = append(numbers, i)
	}
	return numbers
}
