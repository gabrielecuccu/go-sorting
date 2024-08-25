package lib

type SelectionSort struct {
	steps *int
}

func NewSelectionSort() *SelectionSort {
	steps := 0
	return &SelectionSort{
		steps: &steps,
	}
}

func (s SelectionSort) Sort(numbers *[]int) {
	for curSortedKey, curSortedVal := range *numbers {
		var minValKey int
		var minValVal = curSortedVal
		for curUnsortedKey, curUnsortedVal := range (*numbers)[curSortedKey:] {
			*s.steps++
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

func (s SelectionSort) Name() string {
	return "Selection Sort"
}

func (s SelectionSort) Steps() int {
	return *s.steps
}
