package main

import (
	"fmt"

	euclid "github.com/spv-dev/algorithms/recursion/euclid/pkg/euclid_v1"
)

func main() {
	fmt.Println(euclid.Nod(0, 2))
	fmt.Println(euclid.Nod(10, 8))
	fmt.Println(euclid.Nod(30, 12))
	fmt.Println(euclid.Nod(50, 15))
	fmt.Println(euclid.Nod(80, 60))
	fmt.Println(euclid.Nod(10, 10))
	fmt.Println(euclid.Nod(1, 1))

	fmt.Println(euclid.Nod2(0, 0))
	fmt.Println(euclid.Nod2(10, 8))
	fmt.Println(euclid.Nod2(30, 12))
	fmt.Println(euclid.Nod2(50, 15))
	fmt.Println(euclid.Nod2(80, 60))
	fmt.Println(euclid.Nod2(10, 10))
	fmt.Println(euclid.Nod2(1, 1))
}
