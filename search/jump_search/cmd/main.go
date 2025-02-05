package main

import (
	"fmt"

	jump_search "github.com/spv-dev/algorithms/search/jump_search/pkg/jump_search_v1"
)

func main() {
	fmt.Println(jump_search.Find([]int{}, 5))
	fmt.Println(jump_search.Find([]int{1}, 5))
	fmt.Println(jump_search.Find([]int{5}, 5))
	fmt.Println(jump_search.Find([]int{5, 6}, 6))
	fmt.Println(jump_search.Find([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, 0))
	fmt.Println(jump_search.Find([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10))
	fmt.Println(jump_search.Find([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 1))
	fmt.Println(jump_search.Find([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 11))
	fmt.Println(jump_search.Find([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 11))

	fmt.Println(jump_search.Exists([]int{}, 5))
	fmt.Println(jump_search.Exists([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5))
	fmt.Println(jump_search.Exists([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10))
	fmt.Println(jump_search.Exists([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 1))
	fmt.Println(jump_search.Exists([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 11))
}
