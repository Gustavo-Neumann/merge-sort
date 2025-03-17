package main

import "fmt"

func merge(a []int, b []int) []int {
	final := []int{}
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

	for ; i < len(a); i++ {
		final = append(final, a[i])
	}

	for ; j < len(b); j++ {
		final = append(final, b[j])
	}

	return final
}

func MergeSort(items []int) []int {
	if len(items) < 2 {
		return items
	}

	first := MergeSort(items[:len(items)/2])
	second := MergeSort(items[len(items)/2:])

	return merge(first, second)
}

func main() {
	a := []int{10, 2, 8, 1, 3, 4, 6, 7, 9, 5}
	fmt.Println(a)
	fmt.Println(MergeSort(a))
}
