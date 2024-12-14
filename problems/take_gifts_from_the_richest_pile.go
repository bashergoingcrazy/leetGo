// https://leetcode.com/problems/take-gifts-from-the-richest-pile/description/
package problems

import (
	"container/heap"
	"math"
)

type intheap []int

func (h intheap) Len() int {return len(h)}
func (h intheap) Swap(i,j int) {h[i],h[j] = h[j],h[i]}
func (h intheap) Less(i,j int) bool {return h[i] > h[j]}

func (h *intheap) Push(x any) {
    *h = append(*h, x.(int))
}

func (h *intheap) Pop() any {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

func pickGifts(gifts []int, k int) int64 {
    h := &intheap{}
    heap.Init(h)
    for _, v := range gifts {
        heap.Push(h, v)
    }

    for range k {
        top := heap.Pop(h).(int)
        x := int(math.Floor(math.Sqrt(float64(top))))
        heap.Push(h, x)
    }

    count := 0
    for h.Len() > 0 {
        count += heap.Pop(h).(int)
    }
    return int64(count)
}


