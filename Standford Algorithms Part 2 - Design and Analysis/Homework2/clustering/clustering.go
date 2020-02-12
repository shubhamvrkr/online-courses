package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Edge struct {
	startV int
	endV   int
	cost   int
}

func main() {

	var filename = "clustering1.txt"
	edges := readFile(filename)
	sortEdges(edges)
	output := clustering(edges, 4)
	fmt.Println("maximum spacing of a 4-clustering: ", output)
	print()
}

func clustering(edges []Edge, k int) int {
	var edge Edge
	uf, _ := New(500)

	for uf.leaders != k-1 {
		edge = edges[0]
		edges = edges[1:]
		flag, _ := uf.Connected(edge.startV, edge.endV)
		if !flag {
			err := uf.Union(edge.startV, edge.endV)
			if err != nil {
				fmt.Println("Error: ", err.Error())
				os.Exit(1)
			}
		}
	}
	return edge.cost
}

func sortEdges(edges []Edge) {
	sort.Slice(edges, func(i int, j int) bool {
		if edges[i].cost <= edges[j].cost {
			return true
		}
		return false
	})
}

func readFile(filename string) []Edge {

	var edges []Edge = []Edge{}
	file, _ := os.Open(filename)
	defer file.Close()
	fscanner := bufio.NewScanner(file)
	for fscanner.Scan() {
		vertices := strings.Fields(fscanner.Text())
		U, _ := strconv.Atoi(vertices[0])
		V, _ := strconv.Atoi(vertices[1])
		weight, _ := strconv.Atoi(vertices[2])
		edges = append(edges, Edge{U, V, weight})
	}
	return edges
}
