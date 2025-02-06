package interpolationsearch

func Find(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	low := 0
	high := len(nums) - 1

	for (low <= high) && (target >= nums[low]) && (target <= nums[high]) {
		if low == high {
			if nums[low] == target {
				return low
			} else {
				return -1
			}
		}

		//pos := low + (((high - low) / (nums[high] - nums[low])) * (target - nums[low]))

		pos := low + (((target - nums[low]) * (high - low)) / (nums[high] - nums[low]))

		if nums[pos] == target {
			return pos
		}

		if nums[pos] < target {
			low = pos + 1
		} else {
			high = pos - 1
		}
	}

	return -1
}

func Exists(nums []int, target int) bool {
	return Find(nums, target) != -1
}
