package util

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	arr := []int{
		5, 1, 2, 4, 5, 6, 3, 2, 2, 3, 5, 6,
	}

	index := BubbleSort(arr)
	fmt.Println(index)
}
