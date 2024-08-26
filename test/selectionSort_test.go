package test

import (
	"github.com/gabrielecuccu/go-sorting/lib"
	"gotest.tools/v3/assert"
	"reflect"
	"testing"
)

func Test_SelectionSort__SelectionSort_AlgorithmInitialisation(t *testing.T) {
	selectionSort := lib.NewSelectionSort()
	assert.Equal(t, selectionSort.Steps(), 0)
	assert.Equal(t, selectionSort.Name(), "Selection Sort")
}

func Test_SelectionSort_TotallyUnsortedArray(t *testing.T) {
	arr := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	selectionSort := lib.NewSelectionSort()

	selectionSort.Sort(&arr)
	assert.Equal(t, selectionSort.Steps(), 55)
	assert.Assert(t, reflect.DeepEqual([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, arr))
}

func Test_SelectionSort_PartiallyUnsortedArray(t *testing.T) {
	arr := []int{0, 1, 2, 3, 9, 8, 7, 4, 5, 6}
	selectionSort := lib.NewSelectionSort()

	selectionSort.Sort(&arr)
	assert.Equal(t, selectionSort.Steps(), 55)
	assert.Assert(t, reflect.DeepEqual([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, arr))
}

func Test_SelectionSort_SortedArray(t *testing.T) {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	selectionSort := lib.NewSelectionSort()

	selectionSort.Sort(&arr)
	assert.Equal(t, selectionSort.Steps(), 55)
	assert.Assert(t, reflect.DeepEqual([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, arr))
}

func Test_SelectionSort_OneItemArray(t *testing.T) {
	arr := []int{1}
	selectionSort := lib.NewSelectionSort()

	selectionSort.Sort(&arr)
	assert.Equal(t, selectionSort.Steps(), 1)
	assert.Assert(t, reflect.DeepEqual([]int{1}, arr))
}

func Test_SelectionSort_EmptyArray(t *testing.T) {
	arr := []int{}
	selectionSort := lib.NewSelectionSort()

	selectionSort.Sort(&arr)
	assert.Equal(t, selectionSort.Steps(), 0)
	assert.Assert(t, reflect.DeepEqual([]int{}, arr))
}
