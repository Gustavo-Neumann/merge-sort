package main

import (
	"fmt"
	"math/rand"
	"merge-sort/iterative"
	"merge-sort/recursive"
	"os"
	"runtime"
	"time"
)

const arraySize = 10000
const runs = 10

func generateRandomArray(size int) []int {
	arr := make([]int, size)
	for i := range arr {
		arr[i] = rand.Intn(1000000)
	}
	return arr
}

func runBenchmark(threads int) ([]time.Duration, []time.Duration) {
	oldMaxProcs := runtime.GOMAXPROCS(threads)
	defer runtime.GOMAXPROCS(oldMaxProcs)

	iterativeTimes := make([]time.Duration, runs)
	recursiveTimes := make([]time.Duration, runs)

	for i := 0; i < runs; i++ {
		arr := generateRandomArray(arraySize)

		arrCopy := make([]int, arraySize)
		copy(arrCopy, arr)

		start := time.Now()
		iterative.MergeSort(arr)
		iterativeTimes[i] = time.Since(start)

		start = time.Now()
		recursive.MergeSort(arrCopy)
		recursiveTimes[i] = time.Since(start)
	}

	return iterativeTimes, recursiveTimes
}

func main() {
	threadCounts := []int{2, 4, 8, 16}

	iterativeAvgs := make([]float64, len(threadCounts))
	recursiveAvgs := make([]float64, len(threadCounts))

	file, err := os.Create("results.txt")
	if err != nil {
		fmt.Println("Erro ao criar arquivo:", err)
		return
	}
	defer file.Close()

	for i, threads := range threadCounts {
		fmt.Printf("Executando com %d threads...\n", threads)

		iterativeTimes, recursiveTimes := runBenchmark(threads)

		var totalIterative, totalRecursive time.Duration
		for j := 0; j < runs; j++ {
			totalIterative += iterativeTimes[j]
			totalRecursive += recursiveTimes[j]
		}
		avgIterative := totalIterative / time.Duration(runs)
		avgRecursive := totalRecursive / time.Duration(runs)

		iterativeAvgs[i] = avgIterative.Seconds() * 1000
		recursiveAvgs[i] = avgRecursive.Seconds() * 1000

		fmt.Fprintf(file, "Threads: %d\n", threads)

		fmt.Fprintf(file, "Iterative Times (ms):\n")
		for j, t := range iterativeTimes {
			fmt.Fprintf(file, "Run %d: %.3f\n", j+1, t.Seconds()*1000)
		}
		fmt.Fprintf(file, "Avg: %.3f\n\n", avgIterative.Seconds()*1000)

		fmt.Fprintf(file, "Recursive Times (ms):\n")
		for j, t := range recursiveTimes {
			fmt.Fprintf(file, "Run %d: %.3f\n", j+1, t.Seconds()*1000)
		}
		fmt.Fprintf(file, "Avg: %.3f\n\n", avgRecursive.Seconds()*1000)

		fmt.Fprintf(file, "----------------------------\n\n")
	}

	fmt.Fprintf(file, "Threads\tIterative (ms)\tRecursive (ms)\n")
	for i, threads := range threadCounts {
		fmt.Fprintf(file, "%d\t\t\t%.3f\t\t\t\t%.3f\n", threads, iterativeAvgs[i], recursiveAvgs[i])
	}

	fmt.Println("Benchmark concluído! Resultados salvos em 'results.txt'")
}
