package utilities

import "sync"

type Queue struct {
	queue []interface{}
	mutex sync.Mutex
}

func (q *Queue) New() Queue {
	q.Clear()
	return *q
}

func (q *Queue) Clear() {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	q.queue = []interface{}{}
}

func (q *Queue) Pick(index int) interface{} {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	return q.queue[index]
}

func (q *Queue) PickFirst() interface{} {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	return q.queue[0]
}

func (q *Queue) PickEnd() interface{} {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	return q.queue[q.Length()-1]
}

func (q *Queue) Pop(index int) interface{} {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	res := q.queue[index]
	q.queue = append(q.queue[:index], q.queue[index+1:]...)
	return res
}

func (q *Queue) PopFirst() interface{} {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	res := q.queue[0]
	q.queue = q.queue[1:]
	return res
}

func (q *Queue) PushFirst(new ...interface{}) {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	q.queue = append(new, q.queue...)
}

func (q *Queue) PopEnd() interface{} {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	length := q.Length()
	res := q.queue[length-1]
	q.queue = q.queue[:length-1]
	return res
}

func (q *Queue) PushEnd(new ...interface{}) {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	q.queue = append(q.queue, new...)
}

func (q *Queue) PopFifo() interface{} {
	return q.PopFirst()
}

func (q *Queue) PushFifo(new ...interface{}) {
	q.PushEnd(new...)
}

func (q *Queue) PopLifo() interface{} {
	return q.PopEnd()
}

func (q *Queue) PushLifo(new ...interface{}) {
	q.PushEnd(new...)
}

func (q *Queue) Length() int {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	return len(q.queue)
}
