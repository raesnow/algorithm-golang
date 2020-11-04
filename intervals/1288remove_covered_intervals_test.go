package main

import "testing"

func Test_removeCoveredIntervals(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{intervals: [][]int{{1, 4}, {3, 6}, {2, 8}}},
			want: 2,
		},
		{
			name: "test2",
			args: args{intervals: [][]int{{1, 4}, {2, 3}}},
			want: 1,
		},
		{
			name: "test3",
			args: args{intervals: [][]int{{0, 10}, {5, 12}}},
			want: 2,
		},
		{
			name: "test4",
			args: args{intervals: [][]int{{3, 10}, {4, 10}, {5, 11}}},
			want: 2,
		},
		{
			name: "test5",
			args: args{intervals: [][]int{{1, 2}, {1, 4}, {3, 4}}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeCoveredIntervals(tt.args.intervals); got != tt.want {
				t.Errorf("removeCoveredIntervals() = %v, want %v", got, tt.want)
			}
		})
	}
}
