package main

import (
	"bufio"
	"fmt"
	"os"
	"errors"
	"strconv"
	"strings"
)

type Edge struct {
	U int
	V int
	weight int
}

func main() {

	var filename = "graph3.txt"
	edges,v,e := readFile(filename)
	fmt.Println("no of vertices: ",v)
	fmt.Println("no of edges: ",e)
	fmt.Println("length of edges: ",len(edges))

	maxCap,err := AllPairShortest(edges,v,e)
	if err!=nil{
		fmt.Println("Negetive edge cycle exits")
	}else{
			fmt.Println("shortest path: ", maxCap)
	}
}

func AllPairShortest(edges []Edge, V int, E int) (int,error) {

		temp := make([][]int, V+1)
		for i := 1; i <= V; i++ {
   			temp[i] = make([]int, V+1)
		}

		for i:=1;i<=V;i++{
			for j:=1;j<=V;j++{
				if i==j{
					temp[i][j] = 0
				}else{
					temp[i][j] = 9999999
				}
			}
		}

		for _,edge :=range edges {
			temp[edge.U][edge.V] = edge.weight
		}

		for k:=1;k<=V;k++{
			for i:=1;i<=V;i++{
				for j:=1;j<=V;j++{
					if temp[i][k] != 9999999 && temp[k][j] != 9999999 && temp[i][k] + temp[k][j] < temp[i][j]{
						 temp[i][j] = temp[i][k] + temp[k][j]
					}
				}
			}
		}
		var min int = 9999999

		for i:=1;i<=V;i++{
			for j:=1;j<=V;j++{
				if i==j{
					if temp[i][j] <  0 {
						return -1, errors.New("negetive cycle")
					}
				}else{
					if temp[i][j] < min{
						min = temp[i][j]
					}
				}
			}
		}

		return min, nil
}

func readFile(filename string) ([]Edge,int,int) {

	var noOfVertices, noOfEdges int
	var rowCount = 0;
	var edges []Edge = []Edge{}
	file, _ := os.Open(filename)
	defer file.Close()
	fscanner := bufio.NewScanner(file)
	for fscanner.Scan() {
		row := strings.Fields(fscanner.Text())
		if rowCount == 0{
			noOfVertices, _ = strconv.Atoi(row[0])
			noOfEdges, _ = strconv.Atoi(row[1])
			rowCount ++
		}else{
			U, _ := strconv.Atoi(row[0])
			V, _ := strconv.Atoi(row[1])
			weight, _ := strconv.Atoi(row[2])
			edges = append(edges, Edge{U,V,weight})
		}
	}
	return edges,noOfVertices,noOfEdges
}
