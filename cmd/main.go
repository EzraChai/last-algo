package main

import (
	"fmt"

	util "com.github.ezrachai/linear_search"
)

func main() {
	arr := []int{
		1, 2, 3, 4, 5,
	}

	index := util.Search(arr, 2)
	fmt.Println(index)
}
