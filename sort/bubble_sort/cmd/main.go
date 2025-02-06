package main

import (
	"fmt"

	bubblesort "github.com/spv-dev/algorithms/sort/bubble_sort/pkg/bubble_sort_v1"
)

func main() {
	fmt.Println(bubblesort.Sort([]int{}))
	fmt.Println(bubblesort.Sort([]int{7}))
	fmt.Println(bubblesort.Sort([]int{7, 2}))
	fmt.Println(bubblesort.Sort([]int{7, 2, 0}))
	fmt.Println(bubblesort.Sort([]int{1, 7, 2, 0}))
	fmt.Println(bubblesort.Sort([]int{7, 2, 3, 4, 8, 3, 0}))
	fmt.Println(bubblesort.Sort([]int{7, 2, 3, 4, 8, 3, 0, 10, 10, 10}))
	fmt.Println(bubblesort.Sort([]int{1, 2, 3, 4, 5, 6, 7}))
}
