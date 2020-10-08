package main

import "fmt"

func main() {
	intervals := [][]int {
		{2,4},
		{5,7},
		{8,10},
		{11,13},
	}
	newInterval := []int {3,6}
	res := insert(intervals, newInterval)
	fmt.Println(res)
}
