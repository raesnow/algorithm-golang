package main

import "testing"

func Test_findMedianSortedArrays(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "test1",
			args: args{
				nums1: []int{1, 3},
				nums2: []int{2},
			},
			want: 2.0,
		},
		{
			name: "test2",
			args: args{
				nums1: []int{1, 2},
				nums2: []int{3, 4},
			},
			want: 2.5,
		},
		{
			name: "test3",
			args: args{
				nums1: nil,
				nums2: []int{3, 4},
			},
			want: 3.5,
		},
		{
			name: "test4",
			args: args{
				nums1: nil,
				nums2: []int{1},
			},
			want: 1,
		},
		{
			name: "test5",
			args: args{
				nums1: nil,
				nums2: []int{1, 2, 3, 4, 5},
			},
			want: 3,
		},
		{
			name: "test6",
			args: args{
				nums1: []int{1},
				nums2: []int{1},
			},
			want: 1,
		},
		{
			name: "test7",
			args: args{
				nums1: []int{1},
				nums2: []int{2},
			},
			want: 1.5,
		},
		{
			name: "test8",
			args: args{
				nums1: []int{1, 5, 6, 7},
				nums2: []int{2, 3, 4, 8},
			},
			want: 4.5,
		},
		{
			name: "test9",
			args: args{
				nums1: []int{1, 2, 3, 6, 7},
				nums2: []int{4, 5, 8, 9, 10},
			},
			want: 5.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMedianSortedArrays(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("findMedianSortedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}