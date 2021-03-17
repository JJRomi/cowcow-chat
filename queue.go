package main

import (
	sync "sync"
)

// QueueItem 구조체
type QueueItem struct {
	item interface{} // 아이템의 값을 가지고 있습니다.
	prev *QueueItem  // 이전 아이템 구조체를 가지고 있습니다.
}

// Queue 구조체
type Queue struct {
	current *QueueItem // 현재 QueueItem을 가지고 있습니다.
	last    *QueueItem // 가장 마지막 QueueItem을 가지고 있습니다.
	depth   uint64     // 아이템이 몇개가 있는지를 저장합니다.
	//mutex   *sync.Mutex
	mutex *sync.RWMutex
}

// Queue를 생성합니다.
func NewQueue() *Queue {
	var queue *Queue = new(Queue)
	queue.depth = 0
	//queue.mutex = &sync.Mutex{}
	queue.mutex = &sync.RWMutex{}
	return queue
}

// 새로운 아이템을 Queue에 저장합니다.
func (queue *Queue) Enqueue(item interface{}) {
	queue.mutex.Lock()
	defer queue.mutex.Unlock()

	if queue.depth == 0 {
		queue.current = &QueueItem{item: item, prev: nil}
		queue.last = queue.current
		queue.depth++
		return
	}
	q := &QueueItem{item: item, prev: nil}
	queue.last.prev = q
	queue.last = q
	queue.depth++
}

// Queue에서 가장 오래전(처음) 저장된 아이템을 제거합니다.
func (queue *Queue) Dequeue() interface{} {
	queue.mutex.Lock()
	defer queue.mutex.Unlock()
	if queue.depth > 0 {
		item := queue.current.item
		queue.current = queue.current.prev
		queue.depth--
		return item
	}
	return nil
}

func (queue *Queue) Size() uint64 {
	return queue.depth
}
