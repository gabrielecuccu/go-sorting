package test

import (
	"github.com/gabrielecuccu/go-sorting/lib"
	"gotest.tools/v3/assert"
	"reflect"
	"testing"
)

func Test_InsertionSort_AlgorithmInitialisation(t *testing.T) {
	insertionSort := lib.NewInsertionSort()
	assert.Equal(t, insertionSort.Steps(), 0)
	assert.Equal(t, insertionSort.Name(), "Insertion Sort")
}

func Test_InsertionSort_TotallyUnsortedArray(t *testing.T) {
	arr := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	insertionSort := lib.NewInsertionSort()

	insertionSort.Sort(&arr)
	assert.Equal(t, insertionSort.Steps(), 9)
	assert.Assert(t, reflect.DeepEqual([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, arr))
}

func Test_InsertionSort_PartiallyUnsortedArray(t *testing.T) {
	arr := []int{0, 1, 2, 3, 9, 8, 7, 4, 5, 6}
	insertionSort := lib.NewInsertionSort()

	insertionSort.Sort(&arr)
	assert.Equal(t, insertionSort.Steps(), 38)
	assert.Assert(t, reflect.DeepEqual([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, arr))
}

func Test_InsertionSort_SortedArray(t *testing.T) {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	insertionSort := lib.NewInsertionSort()

	insertionSort.Sort(&arr)
	assert.Equal(t, insertionSort.Steps(), 45)
	assert.Assert(t, reflect.DeepEqual([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, arr))
}

func Test_InsertionSort_OneItemArray(t *testing.T) {
	arr := []int{1}
	insertionSort := lib.NewInsertionSort()

	insertionSort.Sort(&arr)
	assert.Equal(t, insertionSort.Steps(), 0)
	assert.Assert(t, reflect.DeepEqual([]int{1}, arr))
}

func Test_InsertionSort_EmptyArray(t *testing.T) {
	arr := []int{}
	insertionSort := lib.NewInsertionSort()

	insertionSort.Sort(&arr)
	assert.Equal(t, insertionSort.Steps(), 0)
	assert.Assert(t, reflect.DeepEqual([]int{}, arr))
}
