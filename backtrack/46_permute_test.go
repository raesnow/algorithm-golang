package main

import (
	"reflect"
	"testing"
)

func Test_permute1(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "test nil",
			args: args{nums: nil},
			want: nil,
		},
		{
			name: "test1",
			args: args{nums: []int{1, 2, 3}},
			want: [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
		},
		{
			name: "test2",
			args: args{nums: []int{0, 1}},
			want: [][]int{{0, 1}, {1, 0}},
		},
		{
			name: "test3",
			args: args{nums: []int{1}},
			want: [][]int{{1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := permute1(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("permute1() = %v, want %v", got, tt.want)
			}
		})
	}
}
