package simplesearch

func Find(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return i
		}
	}

	return -1
}

func Exists(nums []int, target int) bool {
	return Find(nums, target) != -1
}
