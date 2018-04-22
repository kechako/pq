package pq

import (
	"container/heap"
)

type Comparer interface {
	Less(x interface{}) bool
}

const ReorderQueueBufferSize = 1024

type container []Comparer

func (c container) Len() int { return len(c) }

func (c container) Less(i, j int) bool {
	return c[i].Less(c[j])
}

func (c container) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c *container) Push(x interface{}) {
	item := x.(Comparer)
	*c = append(*c, item)
}

func (c *container) Pop() interface{} {
	old := *c
	n := len(old)
	item := old[n-1]
	*c = old[:n-1]
	return item
}

type Queue struct {
	c container
}

func (q *Queue) Push(x Comparer) {
	heap.Push(&q.c, x)
}

func (q *Queue) Pop() Comparer {
	return heap.Pop(&q.c).(Comparer)
}
