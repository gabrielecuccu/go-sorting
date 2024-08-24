package main

import (
	"fmt"
)

func sort(algorithm SortAlgorithm, numbers *[]int) {
	fmt.Println("******************************************")
	fmt.Printf("%s of %v\n", algorithm.name(), *numbers)
	counter := 0
	algorithm.sort(numbers, &counter)
	fmt.Println("Result: ", *numbers)
	fmt.Printf("In %d steps\n", counter)
	fmt.Println("******************************************")
}

func unsortedArray() *[]int {
	return &[]int{9, 8, 7, 6, 5, 4, 3, 2, 1}
}

func main() {
	sort(NewBubbleSort(), unsortedArray())
	sort(NewSelectionSort(), unsortedArray())
}
