package main

import (
	"fmt"
)

func bubbleSort(numbers *[]int, counter *int) {
	swapped := false
	for k, _ := range *numbers {
		*counter++
		if k == len(*numbers)-1 {
			break
		}

		n0 := (*numbers)[k]
		n1 := (*numbers)[k+1]

		if n0 > n1 {
			(*numbers)[k] = n1
			(*numbers)[k+1] = n0
			swapped = true
		}
	}
	if swapped == true {
		bubbleSort(numbers, counter)
	}
}

func selectionSort(numbers *[]int, counter *int) {
	for curSortedKey, curSortedVal := range *numbers {
		var minValKey int
		var minValVal = curSortedVal
		for curUnsortedKey, curUnsortedVal := range (*numbers)[curSortedKey:] {
			*counter++
			if curUnsortedVal < minValVal {
				minValKey = curUnsortedKey
				minValVal = curUnsortedVal
			}
		}
		if minValKey != -1 {
			(*numbers)[curSortedKey] = minValVal
			(*numbers)[minValKey+curSortedKey] = curSortedVal
		}
	}
}

func execBubbleSort() {
	fmt.Println("******************************************")
	fmt.Println("Bubble Sort of 9, 8, 7, 6, 5, 4, 3, 2, 1")
	numbers := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	counter := 0
	bubbleSort(&numbers, &counter)
	fmt.Println("Result: ", numbers)
	fmt.Printf("In %d steps\n", counter)
	fmt.Println("******************************************")
}

func execSelectionSort() {
	fmt.Println("******************************************")
	fmt.Println("Bubble Sort of 9, 8, 7, 6, 5, 4, 3, 2, 1")
	numbers := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	counter := 0
	selectionSort(&numbers, &counter)
	fmt.Println("Result: ", numbers)
	fmt.Printf("In %d steps\n", counter)
	fmt.Println("******************************************")
}

func main() {
	execBubbleSort()
	execSelectionSort()
}
