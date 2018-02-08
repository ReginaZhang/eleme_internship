package queue

import (
	"container/list"
	"github.com/labstack/echo"
)

type Queue struct {
	trigger chan struct{}
	output  chan echo.Map
	stopped bool
	buffer  *list.List
}

func NewQueue() *Queue {
	q := &Queue{
		trigger: make(chan struct{}),
		output:  make(chan echo.Map),
		buffer:  list.New(),
	}

	go q.run()
	return q
}

func (q *Queue) Enqueue(e echo.Map) {
	q.buffer.PushFront(e)
	select {
	case q.trigger <- struct{}{}:
	default:
	}
}

func (q *Queue) Dequeue() <-chan echo.Map {
	return q.output
}

func (q *Queue) Stop() {
	q.stopped = true
	select {
	case q.trigger <- struct{}{}:
	default:
	}
}

func (q *Queue) run() {
	for range q.trigger {
		for e := q.buffer.Back(); e != nil; e = q.buffer.Back() {
			q.output <- e.Value.(echo.Map)
			q.buffer.Remove(e)
		}
		if q.stopped {
			close(q.output)
			return
		}
	}
}
