package recursive

import "sync"

const threshold = 1000

func merge(a []int, b []int) []int {
	final := make([]int, 0, len(a)+len(b))
	i := 0
	j := 0

	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}

	final = append(final, a[i:]...)
	final = append(final, b[j:]...)
	return final
}

func MergeSort(items []int) []int {
	if len(items) <= 1 {
		return items
	}

	// Switch to sequential sort for small slices
	if len(items) < threshold {
		return merge(MergeSort(items[:len(items)/2]), MergeSort(items[len(items)/2:]))
	}

	var wg sync.WaitGroup
	var left, right []int
	wg.Add(2)

	// Parallel recursive sorting
	go func() {
		defer wg.Done()
		left = MergeSort(items[:len(items)/2])
	}()

	go func() {
		defer wg.Done()
		right = MergeSort(items[len(items)/2:])
	}()

	wg.Wait()
	return merge(left, right)
}
