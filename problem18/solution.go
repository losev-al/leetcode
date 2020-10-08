package main

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 || intervals[len(intervals)-1][1] < newInterval[0] {
		return append(intervals, newInterval)
	}
	var result [][] int
	isFinded := false
	for i, interval := range intervals {
		if intersect(interval, newInterval) {
			if isFinded {
				index := len(result) - 1
				result[index][0] = min(result[index][0], interval[0])
				result[index][1] = max(result[index][1], interval[1])
			} else {
				result = append(intervals[:i], []int{min(interval[0], newInterval[0]), max(interval[1], newInterval[1])})
				isFinded = true
			}
		} else  if isFinded {
			result = append(result, intervals[i:]...)
			return result
		} else if (i == 0 || intervals[i-1][1] < newInterval[0]) && interval[0] > newInterval[1] {
			result = append(intervals[:i], append([][]int {newInterval}, intervals[i:]...)...)
			break
		}
	}
	return result
}

func intersect(a []int, b []int) bool {
	if a[0] < b[0] && a[1] > b[0] {
		return true
	}
	if a[0] >= b[0] && a[0] <= b[1] {
		return true
	}
	if a[1] >= b[0] && a[1] <= b[1] {
		return true
	}
	return false
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
