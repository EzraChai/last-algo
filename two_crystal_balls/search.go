package binarysearch

import "math"

// [NNNNNNNNYYYYY]
//
//	length = 13
func TwoCrystalBalls(breaks []bool) int {
	length := len(breaks)
	jump := int(math.Floor(math.Sqrt(float64(length))))

	i := jump
	for ; i < length; i += jump {
		if breaks[i] == true {
			break
		}
	}

	i -= jump

	for j := 0; j < jump && j < length; {
		if breaks[i] == true {
			return i
		}

		j++
		i++
	}

	return -1
}
