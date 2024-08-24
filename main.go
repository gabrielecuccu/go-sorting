package main

import (
	"fmt"
	"github.com/gabrielecuccu/go-sorting/lib"
)

func sort(algorithm lib.SortAlgorithm, numbers *[]int) {
	fmt.Println("******************************************")
	fmt.Printf("%s of %v\n", algorithm.Name(), *numbers)
	counter := 0
	algorithm.Sort(numbers, &counter)
	fmt.Println("Result: ", *numbers)
	fmt.Printf("In %d steps\n", counter)
	fmt.Println("******************************************")
}

func unsortedArray() *[]int {
	return &[]int{9, 8, 7, 6, 5, 4, 3, 2, 1}
}

func main() {
	sort(lib.NewBubbleSort(), unsortedArray())
	sort(lib.NewSelectionSort(), unsortedArray())
}
