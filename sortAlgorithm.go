package main

type SortAlgorithm interface {
	sort(numbers *[]int, counter *int)
	name() string
}
