package main

// Предположим, что эта очередь будет оперировать только положительными
// числами (отрицательные числа ей никогда не поступят на вход)

type numns interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type CircularQueue[T numns] struct {
	start, end int
	queue      []T
}

func NewCircularQueue[T numns](size int) CircularQueue[T] { // создать очередь с определенным размером буффера
	return CircularQueue[T]{
		start: -1,
		end:   -1,
		queue: make([]T, size),
	}
}

func (q *CircularQueue[T]) Push(value T) bool { // добавить значение в конец очереди (false, если очередь заполнена)
	if q.Full() {
		return false
	}

	if q.Empty() {
		q.queue[0] = value
		q.end, q.start = 0, 0

		return true
	}

	nextEnd := q.getNextIndex(q.end)
	q.queue[nextEnd] = value
	q.end = nextEnd

	return true
}

func (q *CircularQueue[T]) Pop() bool { // удалить значение из начала очереди (false, если очередь пустая)
	if q.Empty() {
		return false
	}

	if q.end == q.start {
		q.end, q.start = -1, -1
	}

	q.start = q.getNextIndex(q.start)

	return true
}

func (q *CircularQueue[T]) Front() T { // получить значение из начала очереди (-1, если очередь пустая)
	if q.Empty() {
		return -1
	}

	return q.queue[q.start]
}

func (q *CircularQueue[T]) Back() T { // получить значение из конца очереди (-1, если очередь пустая)
	if q.Empty() {
		return -1
	}

	return q.queue[q.end]
}

func (q *CircularQueue[T]) Empty() bool {
	return q.end == -1
}

func (q *CircularQueue[T]) Full() bool {
	if q.isItZeroQueue() {
		return true
	}

	if q.Empty() {
		return false
	}

	if newEnd := q.getNextIndex(q.end); newEnd == q.start {
		return true
	}

	return false
}

func (q *CircularQueue[T]) getNextIndex(currentIdx int) int {
	return (currentIdx + 1) % len(q.queue)
}

func (q *CircularQueue[T]) isItZeroQueue() bool {
	return (len(q.queue)) == 0
}
