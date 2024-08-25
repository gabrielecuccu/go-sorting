package lib

type InsertionSort struct {
}

func NewInsertionSort() *InsertionSort {
	return &InsertionSort{}
}

func (b InsertionSort) Sort(numbers *[]int, counter *int) {
	for unsortedKey, _ := range (*numbers)[1:] {
		unsortedKey = unsortedKey + 1
		for sortedKey, _ := range (*numbers)[:unsortedKey] {
			*counter++
			if (*numbers)[unsortedKey] < (*numbers)[sortedKey] {
				move(numbers, unsortedKey, sortedKey)
				break
			}
		}
	}
}

func (b InsertionSort) Name() string {
	return "Insertion Sort"
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
