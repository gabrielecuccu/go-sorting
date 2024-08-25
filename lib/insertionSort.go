package lib

type InsertionSort struct {
	steps *int
}

func NewInsertionSort() *InsertionSort {
	steps := 0
	return &InsertionSort{
		steps: &steps,
	}
}

func (i InsertionSort) Sort(numbers *[]int) {
	for unsortedKey, _ := range (*numbers)[1:] {
		unsortedKey = unsortedKey + 1
		for sortedKey, _ := range (*numbers)[:unsortedKey] {
			*i.steps++
			if (*numbers)[unsortedKey] < (*numbers)[sortedKey] {
				move(numbers, unsortedKey, sortedKey)
				break
			}
		}
	}
}

func (i InsertionSort) Name() string {
	return "Insertion Sort"
}

func (i InsertionSort) Steps() int {
	return *i.steps
}

func move(a *[]int, from, to int) {
	elem := (*a)[from]
	remove(a, from)
	insert(a, elem, to)
}

func remove(a *[]int, i int) {
	*a = append((*a)[:i], (*a)[i+1:]...)
}

func insert(a *[]int, c int, i int) {
	*a = append((*a)[:i], append([]int{c}, (*a)[i:]...)...)
}
