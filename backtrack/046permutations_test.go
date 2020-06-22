package main

import (
	"fmt"
	"testing"
)

func Test_permute(t *testing.T) {
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
		{
			name: "test2",
			args: args{nums: []int{5, 4, 6, 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := permute(tt.args.nums)
			println(tt.name)
			for i, v := range got {
				fmt.Printf("%d: %#v\n", i, v)
			}
		})
	}
}
