package main

import (
	"fmt"

	fib "github.com/spv-dev/algorithms/recursion/fibonacci/pkg/fibonacci_v1"
)

func main() {
	fmt.Println(fib.Find(1))
	fmt.Println(fib.Find(2))
	fmt.Println(fib.Find(3))
	fmt.Println(fib.Find(4))
	fmt.Println(fib.Find(5))
	fmt.Println(fib.Find(6))
	fmt.Println(fib.Find(10))
}
