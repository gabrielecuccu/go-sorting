package lib

type SortAlgorithm interface {
	Sort(numbers *[]int, counter *int)
	Name() string
}
