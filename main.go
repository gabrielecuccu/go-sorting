package main

import (
	"fmt"
	"github.com/gabrielecuccu/go-sorting/lib"
)

func sort(algorithm lib.SortAlgorithm, numbers *[]int) {
	fmt.Println("******************************************")
	fmt.Printf("%s of %v\n", algorithm.Name(), *numbers)
	algorithm.Sort(numbers)
	fmt.Println("Result: ", *numbers)
	fmt.Printf("In %d steps\n", algorithm.Steps())
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
	sort(lib.NewBubbleSort(), unsortedArray(25))
	sort(lib.NewSelectionSort(), unsortedArray(25))
	sort(lib.NewInsertionSort(), unsortedArray(25))
	sort(lib.NewMergeSort(), unsortedArray(25))
}
