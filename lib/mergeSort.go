package lib

type MergeSort struct {
	steps *int
}

func NewMergeSort() *MergeSort {
	steps := 0
	return &MergeSort{
		steps: &steps,
	}
}

func (m MergeSort) Sort(numbers *[]int) {
    arrayLength := len(*numbers)
    if arrayLength < 2 {
        return
    }

    middle := arrayLength / 2
    leftArray := (*numbers)[:middle]
    rightArray := (*numbers)[middle:]

    m.Sort(&leftArray)
    m.Sort(&rightArray)

    *m.steps++

    leftCopy := make([]int, len(leftArray))
    copy(leftCopy, leftArray)

    rightCopy := make([]int, len(rightArray))
    copy(rightCopy, rightArray)

    leftArrayKey, rightArrayKey, finalArrayKey := 0, 0, 0

    for leftArrayKey < len(leftCopy) && rightArrayKey < len(rightCopy) {
        if leftCopy[leftArrayKey] < rightCopy[rightArrayKey] {
            (*numbers)[finalArrayKey] = leftCopy[leftArrayKey]
            leftArrayKey++
        } else {
            (*numbers)[finalArrayKey] = rightCopy[rightArrayKey]
            rightArrayKey++
        }
        finalArrayKey++
    }

    for leftArrayKey < len(leftCopy) {
        (*numbers)[finalArrayKey] = leftCopy[leftArrayKey]
        leftArrayKey++
        finalArrayKey++
    }

    for rightArrayKey < len(rightCopy) {
        (*numbers)[finalArrayKey] = rightCopy[rightArrayKey]
        rightArrayKey++
        finalArrayKey++
    }
}

func (m MergeSort) Name() string {
	return "Merge Sort"
}

func (m MergeSort) Steps() int {
	return *m.steps
}
