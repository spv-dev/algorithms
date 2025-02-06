package selectionsort

func Sort(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	copy(res, nums)

	for i := 0; i < n; i++ {
		min := res[i]
		minIndex := i
		for j := i; j < n; j++ {
			if res[j] < min {
				min = res[j]
				minIndex = j
			}
		}
		res[i], res[minIndex] = res[minIndex], res[i]
	}

	return res
}
