package main

import (
	"fmt"

	selectionsort "github.com/spv-dev/algorithms/sort/selection_sort/pkg/selection_sort_v1"
)

func main() {
	fmt.Println(selectionsort.Sort([]int{}))
	fmt.Println(selectionsort.Sort([]int{7}))
	fmt.Println(selectionsort.Sort([]int{7, 2}))
	fmt.Println(selectionsort.Sort([]int{7, 2, 0}))
	fmt.Println(selectionsort.Sort([]int{1, 7, 2, 0}))
	fmt.Println(selectionsort.Sort([]int{7, 2, 3, 4, 8, 3, 0}))
	fmt.Println(selectionsort.Sort([]int{7, 2, 3, 4, 8, 3, 0, 10, 10, 10}))
	fmt.Println(selectionsort.Sort([]int{1, 2, 3, 4, 5, 6, 7}))
}
