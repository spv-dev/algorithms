package main

import (
	"fmt"

	binary_search "github.com/spv-dev/algorithms/search/binary_search/pkg/binary_search_v1"
)

func main() {
	fmt.Println(binary_search.Find([]int{}, 5))
	fmt.Println(binary_search.Find([]int{6}, 6))
	fmt.Println(binary_search.Find([]int{6, 7}, 8))
	fmt.Println(binary_search.Find([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5))
	fmt.Println(binary_search.Find([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10))
	fmt.Println(binary_search.Find([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 1))
	fmt.Println(binary_search.Find([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 11))

	fmt.Println(binary_search.Exists([]int{}, 5))
	fmt.Println(binary_search.Exists([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5))
	fmt.Println(binary_search.Exists([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10))
	fmt.Println(binary_search.Exists([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 1))
	fmt.Println(binary_search.Exists([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 11))
}
