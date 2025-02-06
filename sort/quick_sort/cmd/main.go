package main

import (
	"fmt"

	quicksort "github.com/spv-dev/algorithms/sort/quick_sort/pkg/quick_sort_v1"
)

func main() {
	fmt.Println(quicksort.Sort([]int{}))
	fmt.Println(quicksort.Sort([]int{7}))
	fmt.Println(quicksort.Sort([]int{7, 2}))
	fmt.Println(quicksort.Sort([]int{7, 2, 0}))
	fmt.Println(quicksort.Sort([]int{1, 7, 2, 0}))
	fmt.Println(quicksort.Sort([]int{7, 2, 3, 4, 8, 3, 0}))
	fmt.Println(quicksort.Sort([]int{7, 2, 3, 4, 8, 3, 0, 10, 10, 10}))
	fmt.Println(quicksort.Sort([]int{1, 2, 3, 4, 5, 6, 7}))
	fmt.Println(quicksort.Sort([]int{2, 2, 2, 4, 2, 2, 2}))
}
