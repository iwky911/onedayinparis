package main

import (
	"container/heap"
)

// An Item is something we manage in a priority queue.
type Item struct {
	value    int // The value of the item; arbitrary.
	priority int // The priority of the item in the queue.
	// The index is needed by changePriority and is maintained by the heap.Interface methods.
	index  int // The index of the item in the heap.
	line   string
	parent *Item
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	pq.checkGrow(1)
	a := *pq
	n := len(a)
	a = a[0 : n+1]
	item := x.(*Item)
	item.index = n
	a[n] = item
	*pq = a
}

func (pq *PriorityQueue) checkGrow(g int) {
	l := len(*pq)
	if l+g > cap(*pq) { // reallocate
		// Allocate double what's needed, for future growth.
		newSlice := make([]*Item, l, (l+g)*2)
		// The copy function is predeclared and works for any slice type.
		copy(newSlice, *pq)
		*pq = newSlice
	}
}

func (pq *PriorityQueue) Pop() interface{} {
	a := *pq
	n := len(a)
	item := a[n-1]
	item.index = -1 // for safety
	*pq = a[0 : n-1]
	return item
}

func (pq *PriorityQueue) AddNode(item *Item) {
	heap.Push(pq, item)
}
