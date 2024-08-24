package main

type SelectionSort struct {
}

func NewSelectionSort() *SelectionSort {
	return &SelectionSort{}
}

func (b SelectionSort) sort(numbers *[]int, counter *int) {
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

func (b SelectionSort) name() string {
	return "Selection Sort"
}
