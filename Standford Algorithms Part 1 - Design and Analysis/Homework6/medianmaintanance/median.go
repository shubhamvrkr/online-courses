package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	var filename = "Median.txt"
	count := 0
	array := readIntegerArray(filename)
	var minHeap = NewMinHeap()
	var maxHeap = NewMaxHeap()
	//read element one by one
	for i := 0; i < len(array); i++ {

		//insert
		maxHeap.insert(array[i])

		if !minHeap.isEmpty() && maxHeap.viewMax() > minHeap.viewMin() {
			elem1 := maxHeap.extractMax()
			elem2 := minHeap.extractMin()
			maxHeap.insert(elem2)
			minHeap.insert(elem1)
		}
		//balance
		if maxHeap.getSize() > minHeap.getSize()+1 {
			elem := maxHeap.extractMax()
			minHeap.insert(elem)
		}
		//maxHeap.print()
		//minHeap.print()
		//fmt.Println("Elem: ", maxHeap.viewMax())
		count = (count + maxHeap.viewMax()) % 10000

	}
	fmt.Println("Count: ", count)
}

func readIntegerArray(filename string) []int {
	var list []int
	file, _ := os.Open(filename)
	defer file.Close()
	fscanner := bufio.NewScanner(file)
	for fscanner.Scan() {
		num, _ := strconv.Atoi(fscanner.Text())
		list = append(list, num)
	}
	return list
}
