package util

func BubbleSort(arr []int) []int {

	len := len(arr)
	for k := 0; k < len/2-1; k++ {
		for i := 0; i < len-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i+1], arr[i] = arr[i], arr[i+1]
			}
		}
	}

	return arr
}
