package main

import "testing"

func Test_fib(t *testing.T) {
	type args struct {
		N int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{N: 0},
			want: 0,
		},
		{
			name: "test2",
			args: args{N: 1},
			want: 1,
		},
		{
			name: "test3",
			args: args{N: 2},
			want: 1,
		},
		{
			name: "test4",
			args: args{N: 3},
			want: 2,
		},
		{
			name: "test5",
			args: args{N: 4},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fib(tt.args.N); got != tt.want {
				t.Errorf("fib() = %v, want %v", got, tt.want)
			}
			if got := fib1(tt.args.N); got != tt.want {
				t.Errorf("fib1() = %v, want %v", got, tt.want)
			}
		})
	}
}
