package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	var filename = "test.txt"
	var numbers = readFile(filename)
	cmpCount := countComparison(numbers, 0, len(numbers)-1, 1)
	fmt.Println("count: ", cmpCount)
}

func countComparison(numbers []int, i int, j int, pivotFlag int) int {

	if i >= j {
		return 0
	}

	if pivotFlag == 2 {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	} else if pivotFlag == 3 {
		var midIndex int
		if (j-i+1)%2 == 0 {
			midIndex = (j-i+1)/2 - 1
		} else {
			midIndex = (j - i) / 2
		}
		pivotIndex := getPivotIndex(numbers, i, midIndex, j)
		numbers[i], numbers[pivotIndex] = numbers[pivotIndex], numbers[i]
	}

	c, index := partition(numbers, i, j)
	c1 := countComparison(numbers, i, index-1, pivotFlag)
	c2 := countComparison(numbers, index+1, j, pivotFlag)
	return c + c1 + c2
}

func getPivotIndex(numbers []int, a int, b int, c int) int {
	if numbers[a] > numbers[b] {
		if numbers[b] > numbers[c] {
			return b
		} else if numbers[a] > numbers[c] {
			return c
		} else {
			return a
		}
	} else {
		// Decided a is not greater than b.
		if numbers[a] > numbers[c] {
			return a
		} else if numbers[b] > numbers[c] {
			return c
		} else {
			return b
		}
	}
}

func partition(numbers []int, startIndex int, endIndex int) (int, int) {
	var i, j, pivot int = startIndex + 1, startIndex + 1, numbers[startIndex]
	for j <= endIndex {
		if numbers[j] <= pivot {
			numbers[i], numbers[j] = numbers[j], numbers[i]
			i++
		}
		j++
	}
	numbers[startIndex], numbers[i-1] = numbers[i-1], numbers[startIndex]
	return (endIndex - startIndex), (i - 1)
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
