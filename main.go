package main

import (
	"fmt"
	"github.com/gabrielecuccu/go-sorting/lib"
	"time"
)

func sort(algorithm lib.SortAlgorithm, numbers *[]int) {
	fmt.Println("******************************************")
	fmt.Printf("%s of %v\n", algorithm.Name(), *numbers)
	fmt.Printf("%s\n", algorithm.Name())

	start := float64(time.Now().UnixNano())
	algorithm.Sort(numbers)
	t := float64(time.Now().UnixNano())

	fmt.Println("Result: ", *numbers)
	fmt.Printf("In %d steps, %.3f milliseconds\n", algorithm.Steps(), (t-start)/1e6)
	fmt.Println("******************************************")
}

func unsortedArray(size int) *[]int {
	arr := make([]int, size)
	for k, _ := range arr {
		arr[k] = size - k
	}
	return &arr
}

func main() {
	sort(lib.NewBubbleSort(), unsortedArray(50))
	sort(lib.NewSelectionSort(), unsortedArray(50))
	sort(lib.NewInsertionSort(), unsortedArray(50))
	sort(lib.NewMergeSort(), unsortedArray(50))
}
