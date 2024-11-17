package present_heap

import (
	"container/heap"
	"errors"
)

type Present struct {
	Value int
	Size  int
}

type PresentHeap []Present

func (h PresentHeap) Len() int {
	return len(h)
}

func (h PresentHeap) Less(i, j int) bool {
	if h[i].Value != h[j].Value {
		return h[i].Value > h[j].Value
	}
	return h[i].Size < h[j].Size
}

func (h PresentHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *PresentHeap) Push(x interface{}) {
	*h = append(*h, x.(Present))
}

func (h *PresentHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

func NewPresentHeap(presents []Present) *PresentHeap {
	presentHeap := &PresentHeap{}
	heap.Init(presentHeap)
	for _, present := range presents {
		heap.Push(presentHeap, present)
	}
	return presentHeap
}

func (h *PresentHeap) GetNCoolestPresents(n int) ([]Present, error) {
	if n < 0 || n > h.Len() {
		return nil, errors.New("invalid number of presents requested")
	}

	tempHeap := &PresentHeap{}
	*tempHeap = append(*tempHeap, (*h)...)
	heap.Init(tempHeap)

	coolestPresents := make([]Present, n)
	for i := 0; i < n; i++ {
		coolestPresents[i] = heap.Pop(tempHeap).(Present)
	}

	return coolestPresents, nil
}

type PresentCollection struct {
	presents []Present
}

func NewPresentCollection(presents []Present) *PresentCollection {
	return &PresentCollection{presents: presents}
}

func (pc *PresentCollection) GrabPresents(capacity int) []Present {
	n := len(pc.presents)
	if n == 0 || capacity <= 0 {
		return nil
	}

	dpTable := make([][]int, n+1)
	for i := range dpTable {
		dpTable[i] = make([]int, capacity+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= capacity; j++ {
			if pc.presents[i-1].Size > j {
				dpTable[i][j] = dpTable[i-1][j]
			} else {
				dpTable[i][j] = max(dpTable[i-1][j], pc.presents[i-1].Value+dpTable[i-1][j-pc.presents[i-1].Size])
			}
		}
	}

	var result []Present
	j := capacity
	for i := n; i > 0 && j > 0; i-- {
		if dpTable[i][j] != dpTable[i-1][j] {
			result = append(result, pc.presents[i-1])
			j -= pc.presents[i-1].Size
		}
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
