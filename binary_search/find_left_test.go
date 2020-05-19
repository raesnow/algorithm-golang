package main

import "testing"

func Test_binarySearchLeft(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				nums:   []int{1, 2, 2, 3, 4, 6, 8, 8, 10},
				target: 4,
			},
			want: 4,
		},
		{
			name: "test2",
			args: args{
				nums:   []int{1, 2, 2, 3, 4, 6, 8, 8, 10},
				target: 5,
			},
			want: -1,
		},
		{
			name: "test3",
			args: args{
				nums:   []int{1, 2, 2, 3, 4, 6, 8, 8, 10},
				target: 0,
			},
			want: -1,
		},
		{
			name: "test4",
			args: args{
				nums:   []int{1, 2, 2, 3, 4, 6, 8, 8, 10},
				target: 11,
			},
			want: -1,
		},
		{
			name: "test5",
			args: args{
				nums:   []int{1, 2, 2, 3, 4, 6, 8, 8, 10},
				target: 2,
			},
			want: 1,
		},
		{
			name: "test6",
			args: args{
				nums:   []int{1, 2, 2, 3, 4, 6, 8, 8, 10},
				target: 8,
			},
			want: 6,
		},
		{
			name: "test7",
			args: args{
				nums:   []int{},
				target: 0,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binarySearchLeft1(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("binarySearchLeft1() = %v, want %v", got, tt.want)
			}
		})

		t.Run(tt.name, func(t *testing.T) {
			if got := binarySearchLeft2(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("binarySearchLeft2() = %v, want %v", got, tt.want)
			}
		})
	}
}
