package main

func maxSubArray(nums []int) int {
	result := nums[0]
		sum := 0
	for _, num := range nums {
		sum += num
		if sum>result {
			result = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return result
}
