package queue

type Queue struct {
	items []any
}

func NewQueue() *Queue {
	return &Queue{
		items: make([]any, 0),
	}
}

func (q *Queue) Enqueue(item any) {
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() any {
	top := q.items[0]
	q.items = q.items[1:]
	return top
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}
