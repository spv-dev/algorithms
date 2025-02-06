package main

import (
	"fmt"

	list "github.com/spv-dev/algorithms/recursion/list/pkg/list_v1"
)

func main() {
	fmt.Println(list.Sum([]int{}))
	fmt.Println(list.Sum([]int{1}))
	fmt.Println(list.Sum([]int{1, 2}))
	fmt.Println(list.Sum([]int{1, 2, 3, 4}))
	fmt.Println(list.Sum([]int{1, 2, 3, 4, 5, 6}))
	fmt.Println(list.Sum([]int{1, 2, 3, 4, 5, 6, 10}))
	fmt.Println(list.Sum([]int{1, 2, 3, 4, 5, 6, 100}))

	fmt.Println(list.Length([]int{}))
	fmt.Println(list.Length([]int{1}))
	fmt.Println(list.Length([]int{1, 2}))
	fmt.Println(list.Length([]int{1, 2, 3, 4}))
	fmt.Println(list.Length([]int{1, 2, 3, 4, 5, 6}))
	fmt.Println(list.Length([]int{1, 2, 3, 4, 5, 6, 10}))
	fmt.Println(list.Length([]int{1, 2, 3, 4, 5, 6, 100}))

	fmt.Println(list.Max([]int{}))
	fmt.Println(list.Max([]int{1}))
	fmt.Println(list.Max([]int{1, 2}))
	fmt.Println(list.Max([]int{1, 2, 3, 4}))
	fmt.Println(list.Max([]int{1, 2, 3, 4, 5, 6}))
	fmt.Println(list.Max([]int{1, 2, 30, 4, 5, 6, 10}))
	fmt.Println(list.Max([]int{1, 2, 3, 4, 5, 6, 100}))
}
