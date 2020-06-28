package main

import "testing"

func Test_maxProfit5(t *testing.T) {
	type args struct {
		prices []int
		fee    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				prices: []int{1, 3, 2, 8, 4, 9},
				fee:    2,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit5(tt.args.prices, tt.args.fee); got != tt.want {
				t.Errorf("maxProfit5() = %v, want %v", got, tt.want)
			}
		})
	}
}
