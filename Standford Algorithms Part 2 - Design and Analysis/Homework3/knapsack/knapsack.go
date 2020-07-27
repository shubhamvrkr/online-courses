package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Item struct {
	value int
	weight   int
}

func main() {

	var filename = "knapsack_big.txt"
	items,W,n := readFile(filename)
	fmt.Println("knapsack capacity: ",W)
	fmt.Println("no of items: ",n)
	fmt.Println("length of items: ",len(items))

	maxCap := findMaxCapacity(items,W,n)
	fmt.Println("maximum value: ", maxCap)

}

func findMaxCapacity(items []Item, W int, n int) int {

		temp := make([][]int, n+1)
		for i := 0; i < n+1; i++ {
   			temp[i] = make([]int, W+1)
		}
		for x:=0;x<=W;x++{
			temp[0][x] = 0
		}

		for i:=1;i<=n;i++{
			for x:=0;x<=W;x++ {
					item := items[i-1]
					if(item.weight > x){
						temp[i][x] = temp[i-1][x]
					}else{
						temp[i][x] = max(temp[i-1][x],temp[i-1][x-item.weight] +  item.value )
					}
			}
	}
	return temp[n][W]

}

// OPTIMIZE: traverse items from right and decide max (value(items-lastone, totalweight- item.weight), value(items-lastone,totalweight))
//make decision to include last item or not. use recursion for the same

func max(a int, b int) int{
	if a >b{
		return a
	}
	return b;
}

func readFile(filename string) ([]Item,int,int) {

	var knapsackCapacity, noOfItems int
	var rowCount = 0;
	var items []Item = []Item{}
	file, _ := os.Open(filename)
	defer file.Close()
	fscanner := bufio.NewScanner(file)
	for fscanner.Scan() {
		row := strings.Fields(fscanner.Text())
		if rowCount == 0{
			knapsackCapacity, _ = strconv.Atoi(row[0])
			noOfItems, _ = strconv.Atoi(row[1])
			rowCount ++
		}else{
			value, _ := strconv.Atoi(row[0])
			weight, _ := strconv.Atoi(row[1])
			items = append(items, Item{value,weight})
		}
	}
	return items,knapsackCapacity,noOfItems
}
