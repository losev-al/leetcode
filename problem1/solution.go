package main

func twoSum(nums []int, target int) []int {
	numMap := make(map[int][]int, len(nums))
	for i, num := range nums {
		numMap[num] = append(numMap[num], i)
	}
	for _, num := range nums {
		if secondIndex, ok := numMap[target - num]; ok {
			if num == target - num && len(numMap[target - num]) == 2 {
				return numMap[target - num]
			} else if num != target - num {
				return []int{numMap[num][0], secondIndex[0]}
			}
		}
	}
	return nil
}
