package main

import "fmt"

// heap is a min heap: the min number is at the top.
type heap struct {
	eles []int
}

func New(nums []int) *heap {
	h := &heap{
		make([]int, len(nums)),
	}

	copy(h.eles, nums)

	// for every elements in h.array,
	for i := range h.eles {
		h.heapifyDown(i)
	}

	return h
}

func (h *heap) heapifyUp(idx int) {

}

func (h *heap) heapifyDown(idx int) {
	c, l, r := h.eles[idx], h.eles[leftChild(idx)], h.eles[rightChild(idx)]

	if c < l && c < r {
		return
	} else if c < l {
		c, l = l, c
		h.heapifyDown(leftChild(idx))
	} else {
		c, r = r, c
		h.heapifyDown(rightChild(idx))
	}
}

// leftChild returns left child node's index of the node at heap.eles[idx].
func leftChild(idx int) int {
	return idx*2 + 1
}

// rightChild returns right child node of the node at heap.eles[idx].
func rightChild(idx int) int {
	return idx*2 + 2
}

func (h *heap) String() string {
	return fmt.Sprint(h)
}

func main() {
	h := New([]int{6, 1, 3})
	//should show 1, 3, 6 or 1, 6, 3
	fmt.Println(h.String())
}
