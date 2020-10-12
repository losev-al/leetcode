package main

import (
	"sort"
)

type array2d struct {
	data [][]int
}

func merge(intervals [][]int) [][]int {
	data := array2d{data: intervals}
	sort.Sort(data)
	i := 0
	for i < len(intervals)-1 {
		if intervals[i][1] >= intervals[i+1][0] {
			if intervals[i+1][1] > intervals[i][1] {
				intervals[i][1] = intervals[i+1][1]
			}
			intervals = append(intervals[:i+1], intervals[i+2:]...)
		} else {
			i++
		}
	}
	return intervals
}

func (s array2d) Len() int {
	return len(s.data)
}

func (s array2d) Swap(i, j int) {
	s.data[i], s.data[j] = s.data[j], s.data[i]
}

func (s array2d) Less(i, j int) bool {
	return s.data[i][0] < s.data[j][0]
}
