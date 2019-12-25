// Problem Statement:
// You are given as input an unsorted array of n distinct numbers, where n is a power of 2.
// Give an algorithm that identifies the second-largest number in the array, and that uses at most  n + log(n) - 2 comparisons.

package main

import "fmt"

func main() {

	//Given Facts :
	//N is always power of 2
	//Distinct elements
	var input = []int{5, 3, 18, 24, 12, 89, 13, 38, 199, 200, 90, 28, 94, 27, 14, 49}
	var m map[int][]int = make(map[int][]int)
	var winner, loser int
	var j int = 0
	//calculate the largest number in the array using pair comparison
	//also maintain winners opponent at each level
	//total comparison = n/2 + n/4 + n/8 .... + 1 = n-1 comparisons
	for i := 0; i < len(input)-1; i = i + 2 {
		//compare adjacent elements
		if input[i] > input[i+1] {
			winner = input[i]
			loser = input[i+1]
		} else {
			winner = input[i+1]
			loser = input[i]
		}
		//add oponent to winner list
		defeatedList := m[winner]
		if defeatedList == nil {
			m[winner] = []int{loser}
		} else {
			defeatedList = append(defeatedList, loser)
			m[winner] = defeatedList
		}
		input[j] = winner
		j = j + 1
		if i == len(input)-2 {
			input = input[:(len(input) / 2)]
			i = -2
			j = 0
		}
	}
	opponents := m[input[0]]
	secondLargest := -1
	// Get opponents of winner as second largest number must have competed with the winner
	// total number of opponents competed with winner will be log(n)
	for _, num := range opponents {
		if num > secondLargest {
			secondLargest = num
		}
	}
	fmt.Println("Largest Number:", input[0])
	fmt.Println("Second Largest:", secondLargest)
}
