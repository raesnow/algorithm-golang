package main

import (
	"fmt"
	"testing"
)

func Test_permutation(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test1",
			args: args{nums: []int{1, 2, 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := permutation(tt.args.nums)
			println(tt.name)
			for i, v := range got {
				fmt.Printf("%d: %#v\n", i, v)
			}
		})
	}
}
