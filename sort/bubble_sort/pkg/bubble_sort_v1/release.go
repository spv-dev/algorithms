package bubblesort

func Sort(nums []int) []int {
	n := len(nums)

	result := make([]int, n)
	copy(result, nums)

	exchange := true
	for exchange {
		exchange = false
		for i := 0; i < n-1; i++ {
			if result[i] > result[i+1] {
				result[i], result[i+1] = result[i+1], result[i]
				exchange = true
			}
		}
	}

	return result
}
