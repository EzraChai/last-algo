package binarysearch

func BinarySearch(arr []int, value int) int {
	low := 0
	high := len(arr)
	for low < high {
		mid := (high + low) / 2
		if arr[mid] == value {
			return mid
		} else if arr[mid] > value {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return -1
}
