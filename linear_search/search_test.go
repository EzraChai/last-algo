package util

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	arr := []int{
		1, 2, 3, 4, 5,
	}

	index := Search(arr, 2)
	if index != 1 {
		t.Fatalf(`Search(arr, 2) = %v, want "1", error`, index)
	}
	fmt.Println(index)
}
