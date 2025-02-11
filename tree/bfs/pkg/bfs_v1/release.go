package bfs

import (
	stringqueue "github.com/spv-dev/algorithms/tree/queue/pkg/queue_v1"
	simplequeue "github.com/spv-dev/algorithms/tree/queue/pkg/queue_v2"
)

func HasPath(graph map[string][]string, from string, to string) bool {
	visited := make(map[string]struct{})
	queue := stringqueue.NewStringQueue()
	queue.Enqueue(from)

	for !queue.IsEmpty() {
		item, ok := queue.Dequeue()
		if !ok {
			break
		}

		if _, ok := visited[item]; ok {
			continue
		}

		if item == to {
			return true
		}

		for _, node := range graph[item] {
			queue.Enqueue(node)
		}
	}

	return false
}

type QueueItem struct {
	item string
	path []string
}

func NewQueueItem(item string, parents []string) *QueueItem {
	return &QueueItem{
		item: item,
		path: append(parents, item),
	}
}

func (qi *QueueItem) Item() string {
	return qi.item
}

func (qi *QueueItem) Path() []string {
	return qi.path
}

func GetPath(graph map[string][]string, from string, to string) []string {
	visited := make(map[string]struct{})
	queue := simplequeue.NewQueue()
	queue.Enqueue(NewQueueItem(from, []string{}))

	for !queue.IsEmpty() {
		item := queue.Dequeue().(*QueueItem)

		if _, ok := visited[item.Item()]; ok {
			continue
		}
		visited[item.Item()] = struct{}{}

		if item.Item() == to {
			return item.path
		}

		for _, node := range graph[item.Item()] {
			queue.Enqueue(NewQueueItem(node, item.Path()))
		}
	}

	return []string{}
}
