package binarysearch

func Find(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	low := 0
	high := len(nums) - 1

	for low <= high {
		mid := (low + high) / 2

		elem := nums[mid]

		if elem == target {
			return mid
		}

		if elem > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}

func Exists(nums []int, target int) bool {
	return Find(nums, target) != -1
}
