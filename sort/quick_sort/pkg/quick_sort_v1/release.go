package quicksort

func subArrays(nums []int, pivot int) ([]int, []int) {
	less := make([]int, 0, len(nums)/2)
	greater := make([]int, 0, len(nums)/2)
	for i := 0; i < len(nums); i++ {
		if nums[i] < pivot {
			less = append(less, nums[i])
		} else {
			greater = append(greater, nums[i])
		}
	}

	return less, greater
}

func Sort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	pivot := nums[0]

	less, greater := subArrays(nums[1:], pivot)

	return append(append(Sort(less), pivot), Sort(greater)...)
}
