package main

import (
	"bufio"
	"fmt"
	"os"
	"math"
	"strconv"
	"strings"
)

type City struct {
	id int
	pointA float64
	pointB float64
}

func main() {

	var filename = "tsp_small.txt"
	cities := readFile(filename)
	fmt.Println("no of edges: ",len(cities))
	matrix := distanceMatrix(cities)
	printMatrix(matrix,len(cities))

	minDist := calculateTour(matrix,len(cities))
}

func calculateTour(matrix [][]float64, n int) float64{

	 cache := make(map[string]float64)

	 startCityId = 1;
}

func printMatrix(matrix [][]float64, n int){

	fmt.Println("Distance Matrix: ")
	for i:=1;i<=n;i++{
		for j:=1;j<=n;j++{
			fmt.Print(matrix[i][j]," ")
		}
		fmt.Println()
	}
}

func distanceMatrix(cities []City) [][]float64 {

	matrix := make([][]float64, len(cities)+1)
	for i := range matrix {
	    matrix[i] = make([]float64, len(cities)+1)
	}

	for _,city1 := range cities {
			for _,city2 := range cities {
					if city1.id == city2.id{
						continue;
					}
					matrix[city1.id][city2.id] = distance(city1.pointA, city1.pointB, city2.pointA, city2.pointB)
			}
	}
	return matrix

}

func distance(x1,y1,x2,y2 float64) float64 {

	first := math.Pow(float64(x2-x1), 2)
	second := math.Pow(float64(y2-y1), 2)
	return math.Sqrt(first + second)

}

func readFile(filename string) ([]City) {

	var rowCount = 0;
	var cities []City = []City{}
	file, _ := os.Open(filename)
	defer file.Close()
	fscanner := bufio.NewScanner(file)
	for fscanner.Scan() {
		row := strings.Fields(fscanner.Text())
		if rowCount == 0{
		}else{
			pointA, _ := strconv.ParseFloat(row[0],64)
			pointB, _ := strconv.ParseFloat(row[1],64)
			cities = append(cities, City{rowCount,pointA,pointB})
		}
		rowCount ++
	}
	return cities
}
