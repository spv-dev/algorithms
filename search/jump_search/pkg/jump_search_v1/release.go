package jumpsearch

import (
	"math"
)

func Find(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	if target < nums[0] {
		return -1
	}

	if target > nums[len(nums)-1] {
		return -1
	}

	n := len(nums) - 1
	step := int(math.Sqrt(float64(len(nums))))
	low := 0
	high := 0

	for ; (high <= n) && (nums[high] < target); high += step {
		low = high
	}

	for i := low; i <= high && i <= n; i++ {
		if nums[i] == target {
			return i
		}
	}

	return -1
}

func Exists(nums []int, target int) bool {
	return Find(nums, target) != -1
}
