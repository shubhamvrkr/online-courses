//Problem statement:
// Your task is to code up and run the randomized contraction algorithm for the min cut problem and
// use it on the above graph to compute the min cut.

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {

	minCut := 1000000000000000000
	var filename = "test1.txt"
	_, keys := readFile(filename)
	for i := 1; i < 2*len(keys); i++ {
		rand.Seed(time.Now().UnixNano())
		adjacencyList, keys := readFile(filename)
		count := findMinCut(adjacencyList, keys)
		if count < minCut {
			minCut = count
		}
	}
	fmt.Println("Min Cut: ", minCut)
}

func findMinCut(adjacencyList map[string][]string, keys []string) int {
	minCut := 0
	for len(keys) > 2 {
		//get random U as key from adjacencyList
		U := keys[rand.Intn(len(keys))]
		//get random node from U list from adjacencyList
		val := adjacencyList[U]
		V := val[rand.Intn(len(val))]
		uList := makeSet(append(val, adjacencyList[V]...), U, V)
		adjacencyList[U] = uList
		keys = removeV(keys, V)
		fixAdjacencyList(adjacencyList, V, U)
	}
	for _, v := range adjacencyList {
		minCut = len(v)
		break
	}
	return minCut
}

//remove sink from key set
func removeV(keys []string, v string) []string {
	index := -1
	for i, key := range keys {
		if key == v {
			index = i
			break
		}
	}
	keys[index] = keys[len(keys)-1]
	return keys[:len(keys)-1]
}

//removes self loops
func makeSet(d []string, u string, v string) []string {

	res := make([]string, 0)
	for _, val := range d {
		if val == u || val == v {
			continue
		}
		res = append(res, val)
	}
	return res
}

func fixAdjacencyList(adjacencyList map[string][]string, find string, replace string) {

	//remove V key from adjacencyList as it is merged with U
	delete(adjacencyList, find)

	//replace V with U in other nodes
	for key, value := range adjacencyList {
		adjacencyList[key] = trim(value, find, replace)
	}
}

func trim(tails []string, V string, U string) []string {
	for i, val := range tails {
		if val == V {
			tails[i] = U
		}
	}
	return tails
}

//read adjacencyList from file
func readFile(filename string) (map[string][]string, []string) {
	keys := []string{}
	adjacencyList := make(map[string][]string)
	file, _ := os.Open(filename)
	defer file.Close()
	fscanner := bufio.NewScanner(file)
	for fscanner.Scan() {
		stringArray := strings.Fields(fscanner.Text())
		var key string
		for index, n := range stringArray {
			if index == 0 {
				key = n
				keys = append(keys, key)
				adjacencyList[n] = []string{}
			} else {
				adjacencyList[key] = append(adjacencyList[key], n)
			}
		}
	}

	return adjacencyList, keys
}

//print adjacencyList
func printAdjacencyList(adjacencyList map[string][]string) {
	var keys []string
	for k := range adjacencyList {
		keys = append(keys, k)
	}
	for _, k := range keys {
		fmt.Println(k, " : ", adjacencyList[k])
	}
}
