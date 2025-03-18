package iterative

import (
	"sync"
)

const threshold = 1000

func merge(arr []int, l, m, r int) {
	n1, n2 := m-l+1, r-m
	left, right := make([]int, n1), make([]int, n2)

	copy(left, arr[l:m+1])
	copy(right, arr[m+1:r+1])

	i, j, k := 0, 0, l
	for i < n1 && j < n2 {
		if left[i] <= right[j] {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
		}
		k++
	}

	for i < n1 {
		arr[k] = left[i]
		i++
		k++
	}
	for j < n2 {
		arr[k] = right[j]
		j++
		k++
	}
}

func MergeSort(arr []int) {
	n := len(arr)
	for currSize := 1; currSize < n; currSize *= 2 {
		if currSize < threshold { // This is used for small subarrays
			for leftStart := 0; leftStart < n-1; leftStart += 2 * currSize {
				mid := min(leftStart+currSize-1, n-1)
				rightEnd := min(leftStart+2*currSize-1, n-1)
				merge(arr, leftStart, mid, rightEnd)
			}
		} else {
			var wg sync.WaitGroup
			for leftStart := 0; leftStart < n-1; leftStart += 2 * currSize {
				wg.Add(1)
				go func(ls, cs int) {
					defer wg.Done()
					m := min(ls+cs-1, n-1)
					re := min(ls+2*cs-1, n-1)
					merge(arr, ls, m, re)
				}(leftStart, currSize)
			}
			wg.Wait()
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
