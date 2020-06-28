package main

import "testing"

func Test_maxProfit3(t *testing.T) {
	type args struct {
		k      int
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				k:      2,
				prices: []int{2, 4, 1},
			},
			want: 2,
		},
		{
			name: "test2",
			args: args{
				k:      2,
				prices: []int{3, 2, 6, 5, 0, 3},
			},
			want: 7,
		},
		{
			name: "test3",
			args: args{
				k:      1,
				prices: []int{1, 2},
			},
			want: 1,
		},
		{
			name: "test4",
			args: args{
				k:      2,
				prices: []int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0},
			},
			want: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit3(tt.args.k, tt.args.prices); got != tt.want {
				t.Errorf("maxProfit3() = %v, want %v", got, tt.want)
			}
		})
	}
}
