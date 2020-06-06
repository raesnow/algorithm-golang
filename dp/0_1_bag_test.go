package main

import "testing"

func Test_bag(t *testing.T) {
	type args struct {
		weight  int
		values  []int
		weights []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				weight:  4,
				weights: []int{2, 1, 3},
				values:  []int{4, 2, 3},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bag(tt.args.weight, tt.args.values, tt.args.weights); got != tt.want {
				t.Errorf("bag() = %v, want %v", got, tt.want)
			}
		})
	}
}
