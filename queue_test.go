package pq

import "testing"

type item struct {
	priority int
	name     string
}

func (i *item) Less(x interface{}) bool {
	other := x.(*item)
	return i.priority > other.priority
}

var testItems = []*item{
	{
		priority: 2,
		name:     "alpha",
	},
	{
		priority: 4,
		name:     "bravo",
	},
	{
		priority: 1,
		name:     "charlie",
	},
	{
		priority: 3,
		name:     "delta",
	},
	{
		priority: 5,
		name:     "echo",
	},
}

func Test_Queue(t *testing.T) {
	q := &Queue{}

	for _, i := range testItems {
		q.Push(i)
	}

	for p := 5; p > 0; p-- {
		i := q.Pop().(*item)
		if i.priority != p {
			t.Errorf("Queue.Pop() => item.priority : got %d, want %d", i.priority, p)
		}
	}
}
