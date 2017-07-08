package DTS

import (
	"math"
)

//InsertionSort for array of ints
func InsertionSort(a *[]int) {
	var key, i int
	array := *a
	for j := 1; j < len(array); j++ {
		key = array[j]
		i = j - 1
		for i >= 0 && array[i] > key {
			array[i+1] = array[i]
			i--
		}
		array[i+1] = key
	}
}

//ReverseInsertionSort for array of ints
func ReverseInsertionSort(a *[]int) {
	var key, i int
	array := *a
	length := len(array)
	for j := length - 2; j > -1; j-- {
		key = array[j]
		i = j + 1
		for i < length && array[i] > key {
			array[i-1] = array[i]
			i++
		}
		array[i-1] = key
	}
}

//MaximumSubarray gives the start and end indeces of the maximum subarray aswell as the max int
func MaximumSubarray(a *[]int) (start, end, max int) {
	arr := *a
	max = math.MinInt64
	nstart := 0
	curmax := 0

	for k, v := range arr {
		curmax += v

		if max < curmax {
			max = curmax
			start = nstart
			end = k
		}

		if curmax < 0 {
			curmax = 0
			nstart = k + 1
		}
	}
	return
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
