package main

import (
	"fmt"

	interp_search "github.com/spv-dev/algorithms/search/interpolation_search/pkg/interpolation_search_v1"
)

func main() {
	fmt.Println(interp_search.Find([]int{}, 5))
	fmt.Println(interp_search.Find([]int{1}, 5))
	fmt.Println(interp_search.Find([]int{5}, 5))
	fmt.Println(interp_search.Find([]int{5, 6}, 6))
	fmt.Println(interp_search.Find([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, 0))
	fmt.Println(interp_search.Find([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10))
	fmt.Println(interp_search.Find([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 1))
	fmt.Println(interp_search.Find([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 6))
	fmt.Println(interp_search.Find([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 11))

	fmt.Println(interp_search.Exists([]int{}, 5))
	fmt.Println(interp_search.Exists([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5))
	fmt.Println(interp_search.Exists([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10))
	fmt.Println(interp_search.Exists([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 1))
	fmt.Println(interp_search.Exists([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 11))
}
