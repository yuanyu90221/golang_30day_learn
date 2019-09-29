package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}
func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func main() {
	h := &IntHeap{2, 1, 5}
	heap.Init(h)
	heap.Push(h, 3)
	heap.Push(h, 4)
	heap.Push(h, 9)
	fmt.Printf("minimum: %d\n", (*h)[0])
	heap.Pop(h)
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
	fmt.Println("\nheap:", h)
	h1 := &IntHeap{2, 5}
	heap.Init(h1)
	heap.Push(h1, 3)
	heap.Push(h1, 4)
	heap.Push(h1, 9)
	(*h1)[1] = 6
	// 讓heap重新排序
	heap.Fix(h1, 1)

	for h1.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h1))
	}
	fmt.Printf("\n")
	// 2 4 5 6 9
	privateQueueDemo()
}

type Item struct {
	value    string
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil

	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(item *Item, value string, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

func privateQueueDemo() {
	items := map[string]int{
		"banana": 3, "apple": 2,
		"pear": 4,
	}

	pq := make(PriorityQueue, len(items))

	i := 0

	for value, priority := range items {
		pq[i] = &Item{
			value:    value,
			priority: priority,
			index:    1,
		}
		i++
	}

	heap.Init(&pq)

	// 新增一個新元素
	item := &Item{
		value:    "orange",
		priority: 1,
	}
	heap.Push(&pq, item)
	// 修改該元素
	pq.update(item, item.value, 5)

	// 依序Pop出來
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		fmt.Printf("%.2d:%s ", item.priority, item.value)
	}
	fmt.Println()
}
