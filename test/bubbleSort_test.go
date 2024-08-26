package test

import (
	"github.com/gabrielecuccu/go-sorting/lib"
	"gotest.tools/v3/assert"
	"reflect"
	"testing"
)

func Test_BubbleSort_AlgorithmInitialisation(t *testing.T) {
	bubbleSort := lib.NewBubbleSort()
	assert.Equal(t, bubbleSort.Steps(), 0)
	assert.Equal(t, bubbleSort.Name(), "Bubble Sort")
}

func Test_BubbleSort_TotallyUnsortedArray(t *testing.T) {
	arr := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	bubbleSort := lib.NewBubbleSort()

	bubbleSort.Sort(&arr)
	assert.Equal(t, bubbleSort.Steps(), 100)
	assert.Assert(t, reflect.DeepEqual([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, arr))
}

func Test_BubbleSort_PartiallyUnsortedArray(t *testing.T) {
	arr := []int{0, 1, 2, 3, 9, 8, 7, 4, 5, 6}
	bubbleSort := lib.NewBubbleSort()

	bubbleSort.Sort(&arr)
	assert.Equal(t, bubbleSort.Steps(), 40)
	assert.Assert(t, reflect.DeepEqual([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, arr))
}

func Test_BubbleSort_SortedArray(t *testing.T) {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	bubbleSort := lib.NewBubbleSort()

	bubbleSort.Sort(&arr)
	assert.Equal(t, bubbleSort.Steps(), 10)
	assert.Assert(t, reflect.DeepEqual([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, arr))
}

func Test_BubbleSort_OneItemArray(t *testing.T) {
	arr := []int{1}
	bubbleSort := lib.NewBubbleSort()

	bubbleSort.Sort(&arr)
	assert.Equal(t, bubbleSort.Steps(), 1)
	assert.Assert(t, reflect.DeepEqual([]int{1}, arr))
}

func Test_BubbleSort_EmptyArray(t *testing.T) {
	arr := []int{}
	bubbleSort := lib.NewBubbleSort()

	bubbleSort.Sort(&arr)
	assert.Equal(t, bubbleSort.Steps(), 0)
	assert.Assert(t, reflect.DeepEqual([]int{}, arr))
}
