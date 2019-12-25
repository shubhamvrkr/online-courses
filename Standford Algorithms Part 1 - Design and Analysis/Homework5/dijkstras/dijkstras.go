package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Edge struct {
	label  string
	weight int
}

func main() {

	var filename = "dijkstraData.txt"
	adjacencyList := readAssignmentFile(filename)
	output := computeShortestPath(adjacencyList, "1")
	for key, value := range output {
		fmt.Println(key, " ", value)
	}
	fmt.Println()
}

func computeShortestPath(graph map[string][]Edge, sourceV string) map[string]int {
	output := make(map[string]int)
	minHeap := NewMinHeap()
	for key, _ := range graph {
		minHeap.insert(Node{key, MAX_INFINITY})
	}
	minHeap.decreaseKey(sourceV, 0)
	for !minHeap.isEmpty() {
		minNode := minHeap.extractMin()
		for _, edge := range graph[minNode.label] {
			if minHeap.getNodeByLabel(edge.label) != nil && minNode.key+edge.weight < minHeap.getNodeByLabel(edge.label).key {
				minHeap.decreaseKey(edge.label, minNode.key+edge.weight)
			}
		}
		output[minNode.label] = minNode.key
	}
	return output
}

func printAdjacencyList(adjacencyList map[string][]Edge) {
	var keys []string
	for k := range adjacencyList {
		keys = append(keys, k)
	}
	for _, k := range keys {
		fmt.Println(k, " : ", adjacencyList[k])
	}
}

func readFile(filename string) map[string][]Edge {
	var adjG map[string][]Edge = make(map[string][]Edge)
	file, _ := os.Open(filename)
	defer file.Close()
	fscanner := bufio.NewScanner(file)
	for fscanner.Scan() {
		vertices := strings.Fields(fscanner.Text())
		U := vertices[0]
		V := vertices[1]
		weight, _ := strconv.Atoi(vertices[2])
		if adjG[V] == nil {
			adjG[V] = []Edge{}
		}

		if adjG[U] == nil {
			list := []Edge{}
			list = append(list, Edge{V, weight})
			adjG[U] = list
		} else {
			list := adjG[U]
			list = append(list, Edge{V, weight})
			adjG[U] = list
		}
	}
	return adjG
}

func readAssignmentFile(filename string) map[string][]Edge {
	var adjG map[string][]Edge = make(map[string][]Edge)
	file, _ := os.Open(filename)
	defer file.Close()
	fscanner := bufio.NewScanner(file)
	for fscanner.Scan() {
		vertices := strings.Fields(fscanner.Text())
		U := vertices[0]
		tempList := vertices[1:len(vertices)]
		if adjG[U] == nil {
			adjG[U] = []Edge{}
		}
		for _, item := range tempList {
			edge := strings.Split(item, ",")
			weight, _ := strconv.Atoi(edge[1])
			edges := adjG[U]
			edges = append(edges, Edge{edge[0], weight})
			adjG[U] = edges
		}
	}
	return adjG
}
