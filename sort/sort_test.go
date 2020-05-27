package sort

import (
<<<<<<< HEAD
=======
	"reflect"
>>>>>>> fb02a88f07d2997b305b92f22e37d19802aba3a2
	"testing"
)

func Test_mergeSort(t *testing.T) {
	got := mergeSort([]int{7, 6, 5, 4, 2, 9})
	t.Log(got)
}

<<<<<<< HEAD
func Test_heapSort(t *testing.T) {
	got := heapSort([]int{7, 6, 5, 4, 2, 9})
	t.Log(got)
}

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
=======
func Test_mergeSort1(t *testing.T) {
	type args struct {
		nums []int
>>>>>>> fb02a88f07d2997b305b92f22e37d19802aba3a2
	}
	tests := []struct {
		name string
		args args
<<<<<<< HEAD
		want int
	}{
		{
			name: "das",
			args: args{
				nums: []int{
					4, 5, 6, 7, 0, 1, 2,
				},
				target: 0,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
=======
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeSort(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
>>>>>>> fb02a88f07d2997b305b92f22e37d19802aba3a2
