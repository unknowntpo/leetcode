package main

import "fmt"

// if in is nil, return nil.
func mergeSort(in []int) []int {
	// if len(in) <= 1
	// // return in

	if len(in) <= 1 {
		return in
	}
	// find the mid
	mid := len(in) / 2
	// left = mergeSort(in[:mid])
	left := mergeSort(in[:mid])
	// right = mergeSort(in[mid:])
	right := mergeSort(in[mid:])

	// return merge(left, right)
	return merge(left, right)
}

func merge(l, r []int) []int {
	// fast path: return if no need to sort.
	//if len(l) == 0; return r
	if len(l) == 0 {
		return r
	}
	//if len(r) == 0; return l
	if len(r) == 0 {
		return l
	}
	// TODO: in-place swap
	out := []int{}
	idxL, idxR := 0, 0
	for {
		if l[idxL] < r[idxR] {
			out = append(out, l[idxL])
			idxL++
			if idxL >= len(l) {
				out = append(out, r...)
				break
			}
		} else {
			out = append(out, r[idxR])
			idxR++
			if idxR >= len(r) {
				out = append(out, l...)
				break
			}
		}
	}
	return out
}

func main() {
	in := []int{5, 4, 3, 2, 1}
	in2 := []int{}
	in3 := []int{5, 5, 5, 1, 1}

	fmt.Println(mergeSort(in))
	fmt.Println(mergeSort(in2))

	fmt.Println(mergeSort(in3))
}
