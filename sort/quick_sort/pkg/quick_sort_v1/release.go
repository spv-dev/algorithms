package quicksort

import (
	"math/rand"
)

func subArrays(nums []int, pivotIndex int) ([]int, []int) {
	pivot := nums[pivotIndex]

	less := make([]int, 0, len(nums)/2)
	greater := make([]int, 0, len(nums)/2)
	for i := 0; i < len(nums); i++ {
		if i != pivotIndex {
			if nums[i] < pivot {
				less = append(less, nums[i])
			} else {
				greater = append(greater, nums[i])
			}
		}
	}

	return less, greater
}

func Sort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	if len(nums) == 2 {
		if nums[0] > nums[1] {
			return []int{nums[1], nums[0]}
		} else {
			return nums
		}
	}

	pivotIndex := rand.Intn(len(nums))

	less, greater := subArrays(nums, pivotIndex)

	return append(append(Sort(less), nums[pivotIndex]), Sort(greater)...)
}
