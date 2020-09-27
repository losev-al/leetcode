package problem11

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type testData struct {
	expected [][]int
	candidates []int
	target int
}

func TestCombinationSum(t *testing.T) {
	data := []testData {
		{
			expected: [][]int{
				{7},
				{2, 2, 3},
			},
			candidates: []int{2, 3, 6, 7},
			target:     7,
		},
		{
			expected: [][]int{
				{2,2,2,2},
				{2,3,3},
				{3,5},
			},
			candidates: []int{2,3,5},
			target:     8,
		},
	}
	for _, datum := range data {
		actual := combinationSum(datum.candidates, datum.target)
		assert.Equal(t, len(actual), len(datum.expected))
		for _, exp := range datum.expected {
			isOk := false
			for _, act := range actual {
				if len(exp) == len(act) {
					isEq := true
					for i, _ := range exp {
						if exp[i] != act[i] {
							isEq = false
							break
						}
					}
					if isEq {
						isOk = true
						break
					}
				}
			}
			assert.True(t, isOk)
		}
	}
}
