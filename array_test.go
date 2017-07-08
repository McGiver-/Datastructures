package DTS

import "testing"
import "fmt"

func TestInsertionSort(t *testing.T) {
	type args struct {
		a *[]int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"insertionSort",
			args{
				&[]int{3, 10, 2, 1, 25, 22, 111},
			},
		},
		{
			"insertionSort",
			args{
				&[]int{5, 4, 3, 2, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InsertionSort(tt.args.a)
			fmt.Printf("array sorted %v\n", tt.args.a)
		})
	}
}

func TestReverseInsertionSort(t *testing.T) {
	type args struct {
		a *[]int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"insertionSort",
			args{
				&[]int{3, 10, 2, 1, 25, 22, 111},
			},
		},
		{
			"insertionSort",
			args{
				&[]int{5, 4, 3, 2, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReverseInsertionSort(tt.args.a)
			fmt.Printf("array sorted decresing %v\n", tt.args.a)
		})
	}
}
func TestMaximumSubarray(t *testing.T) {
	type args struct {
		a *[]int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"maximumSubArray",
			args{
				&[]int{-2, -3, 4, -1, -2, 1, 5, -3},
			},
		},
		{
			"maximumSubarray",
			args{
				&[]int{-1, -4, 2, 5, 6, -3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start, end, max := MaximumSubarray(tt.args.a)
			fmt.Printf("start %v end %v max %v\n", start, end, max)
		})
	}
}
