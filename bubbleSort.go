package main

type BubbleSort struct {
}

func NewBubbleSort() *BubbleSort {
	return &BubbleSort{}
}

func (b BubbleSort) sort(numbers *[]int, counter *int) {
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
		b.sort(numbers, counter)
	}
}

func (b BubbleSort) name() string {
	return "Bubble Sort"
}
