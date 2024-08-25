package lib

type SortAlgorithm interface {
	Sort(numbers *[]int)
	Name() string
	Steps() int
}
