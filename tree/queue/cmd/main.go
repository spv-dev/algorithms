package main

import (
	"fmt"

	queue "github.com/spv-dev/algorithms/tree/queue/pkg/queue_v1"
)

func main() {
	q := queue.NewStringQueue()
	q.Enqueue("Mike")
	q.Enqueue("Tom")
	q.Enqueue("Bob")
	top, ok := q.Top()
	fmt.Println("Top = :", top)
	fmt.Println("IsEmpty = :", ok)
	val, _ := q.Dequeue()
	fmt.Println("Deque = :", val)
	val, _ = q.Dequeue()
	fmt.Println("Deque = :", val)
	val, _ = q.Dequeue()
	fmt.Println("Deque = :", val)
	fmt.Println("IsEmpty = :", ok)
	val, _ = q.Dequeue()
	fmt.Println("Deque = :", val)
}
