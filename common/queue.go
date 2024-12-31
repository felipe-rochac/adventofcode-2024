package common

// Queue type
type Queue[T any] struct {
	items []T
}

// Enqueue adds an item to the end of the queue
func (q *Queue[T]) Push(item T) {
	q.items = append(q.items, item)
}

// Dequeue removes and returns the item from the front of the queue
func (q *Queue[T]) Pop() (T, bool) {
	if len(q.items) == 0 {
		var zero T
		return zero, false
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, true
}

// IsEmpty checks if the queue is empty
func (q *Queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}

// Size returns the number of items in the queue
func (q *Queue[T]) Size() int {
	return len(q.items)
}

func (q *Queue[T]) Empty() {
	q.items = make([]T, 0)
}
