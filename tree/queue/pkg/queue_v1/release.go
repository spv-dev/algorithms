package queue

type StringQueue struct {
	queue []string
}

func NewStringQueue() *StringQueue {
	return &StringQueue{
		queue: make([]string, 0),
	}
}

func (q *StringQueue) Enqueue(s string) {
	q.queue = append(q.queue, s)
}

func (q *StringQueue) Top() (string, bool) {
	if q.IsEmpty() {
		return "", false
	}
	return q.queue[0], true
}

func (q *StringQueue) IsEmpty() bool {
	return len(q.queue) == 0
}

func (q *StringQueue) Dequeue() (string, bool) {
	result, ok := q.Top()
	if ok {
		q.queue = q.queue[1:]
	}
	return result, ok
}
