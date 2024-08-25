package binarysearch

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	arr := []bool{
		false, false, false, false, false, false, false, false, false, false, true, true, true, true, true,
	}

	index := TwoCrystalBalls(arr)

	fmt.Println(index)
}
