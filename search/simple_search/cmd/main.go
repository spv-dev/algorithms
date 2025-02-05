package main

import (
	"fmt"

	simple_search "github.com/spv-dev/algorithms/search/simple_search/pkg/simple_search_v1"
)

func main() {
	fmt.Println(simple_search.Find([]int{}, 5))
	fmt.Println(simple_search.Find([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5))
	fmt.Println(simple_search.Find([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10))
	fmt.Println(simple_search.Find([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 1))
	fmt.Println(simple_search.Find([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 11))

	fmt.Println(simple_search.Exists([]int{}, 5))
	fmt.Println(simple_search.Exists([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5))
	fmt.Println(simple_search.Exists([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10))
	fmt.Println(simple_search.Exists([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 1))
	fmt.Println(simple_search.Exists([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 11))
}
