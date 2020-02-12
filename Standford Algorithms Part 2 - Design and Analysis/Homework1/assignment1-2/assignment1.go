package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Job struct {
	weight int
	length int
	diff   int
	id     int
}

func main() {

	sumLength := 0
	completionTime := 0

	fileName := "jobs.txt"

	jobs, _ := readFile(fileName)

	//sortByDiff(jobs)
	sortByRatio(jobs)

	for _, job := range jobs {
		sumLength = sumLength + job.length
		completionTime += job.weight * sumLength
	}

	fmt.Println("Final Completion Time: ", completionTime)
}

func sortByDiff(jobs []Job) {
	sort.Slice(jobs, func(i int, j int) bool {
		if jobs[i].diff > jobs[j].diff {
			return true
		} else if jobs[i].diff == jobs[j].diff {
			if jobs[i].weight > jobs[j].weight {
				return true
			}
		}
		return false
	})
}

func sortByRatio(jobs []Job) {
	sort.Slice(jobs, func(i int, j int) bool {
		if float64(jobs[i].weight)/float64(jobs[i].length) > float64(jobs[j].weight)/float64(jobs[j].length) {
			return true
		}
		return false
	})
}

func readFile(filename string) ([]Job, int) {
	index := 0
	count := 0
	var jobs []Job

	file, _ := os.Open(filename)
	defer file.Close()
	fscanner := bufio.NewScanner(file)
	for fscanner.Scan() {
		if index == 0 {
			count, _ = strconv.Atoi(fscanner.Text())
		} else {
			labels := strings.Fields(fscanner.Text())
			w, _ := strconv.Atoi(labels[0])
			l, _ := strconv.Atoi(labels[1])
			job := Job{w, l, w - l, index}
			jobs = append(jobs, job)
		}
		index++
	}
	return jobs, count
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
