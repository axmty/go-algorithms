package sorts

import "fmt"

// Bubble sorts arr using bubble sort algorithm.
func Bubble(arr []int) {
	var swapped bool = true
	for n := len(arr) - 1; swapped; n-- {
		swapped = false
		for j := 0; j < n; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
	}
}

// Insertion sorts arr using insertion sort algorithm.
func Insertion(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i; j > 0 && arr[j-1] > arr[j]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
}

// Merge sorts arr using merge sort algorithm.
func Merge(arr []int) {
	var rec func(int, int)
	rec = func(begin, end int) {
		if end-begin < 2 {
			return
		}
		mid := (begin + end) / 2
		rec(begin, mid)
		rec(mid, end)
		sorted := make([]int, end-begin)
		for i, j := begin, mid; i < mid || j < end; {
			fmt.Println(i, j)
			if j >= end || (i < mid && arr[i] < arr[j]) {
				sorted[i+j-begin-mid] = arr[i]
				i++
			} else {
				sorted[i+j-begin-mid] = arr[j]
				j++
			}
		}
		copySlice(sorted, arr, begin)
	}
	rec(0, len(arr))
	fmt.Println(arr)
}

// Selection sorts arr using selection sort algorithm.
func Selection(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

func copySlice(src, dest []int, destBegin int) {
	for i, v := range src {
		dest[destBegin+i] = v
	}
}
