package iterative

func merge(arr []int, l, m, r int) {
	n1, n2 := m-l+1, r-m
	left, right := make([]int, n1), make([]int, n2)

	for i := 0; i < n1; i++ {
		left[i] = arr[l+i]
	}
	for j := 0; j < n2; j++ {
		right[j] = arr[m+1+j]
	}

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
		for leftStart := 0; leftStart < n-1; leftStart += 2 * currSize {
			mid := min(leftStart+currSize-1, n-1)
			rightEnd := min(leftStart+2*currSize-1, n-1)
			merge(arr, leftStart, mid, rightEnd)
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
