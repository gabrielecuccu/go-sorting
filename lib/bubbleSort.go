package lib

type BubbleSort struct {
	steps *int
}

func NewBubbleSort() *BubbleSort {
	steps := 0
	return &BubbleSort{
		steps: &steps,
	}
}

func (b BubbleSort) Sort(numbers *[]int) {
	swapped := false
	for k, _ := range *numbers {
		*b.steps++
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
		b.Sort(numbers)
	}
}

func (b BubbleSort) Name() string {
	return "Bubble Sort"
}

func (b BubbleSort) Steps() int {
	return *b.steps
}
