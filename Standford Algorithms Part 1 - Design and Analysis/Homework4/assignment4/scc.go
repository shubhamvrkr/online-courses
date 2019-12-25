//Problem statement:
// Your task is to code up and run the randomized contraction algorithm for the min cut problem and
// use it on the above graph to compute the min cut.

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/golang-collections/collections/stack"
)

func main() {
	sccSizes := []int{}
	s := stack.New()
	var filename = "SCC.txt"
	adjacenctyList, reversedAdjList := readFile(filename)
	isVisitedA, isVisitedB := make(map[string]bool), make(map[string]bool)
	for key, _ := range adjacenctyList {
		isVisitedA[key] = false
		isVisitedB[key] = false
	}
	fmt.Println(len(isVisitedA))
	for key, _ := range reversedAdjList {
		if isVisitedA[key] == false {
			recordFinishTime(reversedAdjList, s, key, isVisitedA)
		}
	}
	fmt.Println("Stack Length: ", s.Len())
	for s.Peek() != nil {
		k := fmt.Sprintf("%v", s.Pop())
		if isVisitedB[k] == false {
			scc := findSCC(adjacenctyList, k, isVisitedB)
			sccSizes = append(sccSizes, len(scc))
		}
	}
	sort.Ints(sccSizes)
	fmt.Println("Sorted: ", sccSizes[len(sccSizes)-10:len(sccSizes)])
}

func findSCC(adjList map[string][]string, startVertex string, isVisited map[string]bool) []string {
	list := []string{}
	if isVisited[startVertex] == true {
		return list
	}
	isVisited[startVertex] = true
	for _, nextV := range adjList[startVertex] {
		list = append(list, findSCC(adjList, nextV, isVisited)...)
	}
	return append(list, startVertex)
}

func recordFinishTime(adjList map[string][]string, s *stack.Stack, startVertex string, isVisited map[string]bool) {
	if isVisited[startVertex] == true {
		return
	}
	isVisited[startVertex] = true
	list := adjList[startVertex]
	for _, nextV := range list {
		recordFinishTime(adjList, s, nextV, isVisited)
	}
	s.Push(startVertex)
}

func readFile(filename string) (map[string][]string, map[string][]string) {
	var adjG map[string][]string = make(map[string][]string)
	var reverseAdjG map[string][]string = make(map[string][]string)

	file, _ := os.Open(filename)
	defer file.Close()
	fscanner := bufio.NewScanner(file)
	for fscanner.Scan() {
		vertices := strings.Fields(fscanner.Text())
		U := vertices[0]
		V := vertices[1]

		if adjG[V] == nil {
			adjG[V] = []string{}
		}
		if reverseAdjG[U] == nil {
			reverseAdjG[U] = []string{}
		}
		if adjG[U] == nil {
			list := []string{}
			list = append(list, V)
			adjG[U] = list
		} else {
			list := adjG[U]
			list = append(list, V)
			adjG[U] = list
		}
		if reverseAdjG[V] == nil {
			list := []string{}
			list = append(list, U)
			reverseAdjG[V] = list
		} else {
			list := reverseAdjG[V]
			list = append(list, U)
			reverseAdjG[V] = list
		}
	}
	return adjG, reverseAdjG
}

func printAdjacencyList(adjacencyList map[string][]string) {
	var keys []string
	for k := range adjacencyList {
		keys = append(keys, k)
	}
	for _, k := range keys {
		fmt.Println(k, " : ", adjacencyList[k])
	}
}
