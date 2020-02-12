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

var set map[string]bool = map[string]bool{}

func main() {

	var filename = "edges.txt"
	adjacencyList := readFile(filename)
	output := primsAlgorithm(adjacencyList, "1")
	fmt.Println("Minimum Cost Spanning Tree: ", output)
}

func primsAlgorithm(graph map[string][]Edge, sourceV string) int {

	minCost := 0
	minHeap := NewMinHeap()
	//pick random source
	for _, edge := range graph[sourceV] {
		minHeap.insert(Node{edge.label, edge.weight})
	}
	//add the first vertex
	addVertexToSet(sourceV)
	//condition will work only for connected graphs, so assume graph is connected
	for len(set) < len(graph) {

		minEdge := minHeap.extractMin()
		if isVertexAlreadyInSet(minEdge.label) {
			continue
		}
		fmt.Println("Adding edge ", minEdge.label+" with cost: ", minEdge.key)
		minCost += minEdge.key
		addVertexToSet(minEdge.label)
		for _, edge := range graph[minEdge.label] {
			if !isVertexAlreadyInSet(edge.label) {
				minHeap.insert(Node{edge.label, edge.weight})
			}
		}
	}
	return minCost
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
		//add edge to both vertices adjacencyList
		if adjG[V] == nil {
			list := []Edge{}
			list = append(list, Edge{U, weight})
			adjG[V] = list

		} else {
			list := adjG[V]
			list = append(list, Edge{U, weight})
			adjG[V] = list
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

func addVertexToSet(vertex string) {
	set[vertex] = true
}

func isVertexAlreadyInSet(vertex string) bool {
	return set[vertex]
}

func getEdgeLabel(sourceV, endV string) string {
	if sourceV < endV {
		return sourceV + "-" + endV
	}
	return endV + "-" + sourceV
}
