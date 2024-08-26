package test

import (
	"github.com/gabrielecuccu/go-sorting/lib"
	"gotest.tools/v3/assert"
	"reflect"
	"testing"
)

func Test_MergeSort_AlgorithmInitialisation(t *testing.T) {
	mergeSort := lib.NewMergeSort()
	assert.Equal(t, mergeSort.Steps(), 0)
	assert.Equal(t, mergeSort.Name(), "Merge Sort")
}

func Test_MergeSort_TotallyUnsortedArray(t *testing.T) {
	arr := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	mergeSort := lib.NewMergeSort()

	mergeSort.Sort(&arr)
	assert.Equal(t, mergeSort.Steps(), 9)
	assert.Assert(t, reflect.DeepEqual([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, arr))
}

func Test_MergeSort_PartiallyUnsortedArray(t *testing.T) {
	arr := []int{0, 1, 2, 3, 9, 8, 7, 4, 5, 6}
	mergeSort := lib.NewMergeSort()

	mergeSort.Sort(&arr)
	assert.Equal(t, mergeSort.Steps(), 9)
	assert.Assert(t, reflect.DeepEqual([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, arr))
}

func Test_MergeSort_SortedArray(t *testing.T) {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	mergeSort := lib.NewMergeSort()

	mergeSort.Sort(&arr)
	assert.Equal(t, mergeSort.Steps(), 9)
	assert.Assert(t, reflect.DeepEqual([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, arr))
}

func Test_MergeSort_OneItemArray(t *testing.T) {
	arr := []int{1}
	mergeSort := lib.NewMergeSort()

	mergeSort.Sort(&arr)
	assert.Equal(t, mergeSort.Steps(), 0)
	assert.Assert(t, reflect.DeepEqual([]int{1}, arr))
}

func Test_MergeSort_EmptyArray(t *testing.T) {
	arr := []int{}
	mergeSort := lib.NewMergeSort()

	mergeSort.Sort(&arr)
	assert.Equal(t, mergeSort.Steps(), 0)
	assert.Assert(t, reflect.DeepEqual([]int{}, arr))
}
