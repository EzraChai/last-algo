package binarysearch

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	arr := []int{
		1, 2, 3, 4, 5,
	}

	index := BinarySearch(arr, 2)
	if index != 1 {
		t.Fatalf(`BinarySearch(arr, 2) = %v, want "1", error`, index)
	}
	fmt.Println(index)
}
