package main

import "testing"

func Test_fib2(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{n: 1},
			want: 1,
		},
		{
			name: "test2",
			args: args{n: 2},
			want: 1,
		},
		{
			name: "test3",
			args: args{n: 3},
			want: 2,
		},
		{
			name: "test4",
			args: args{n: 4},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fib2(tt.args.n); got != tt.want {
				t.Errorf("fib2() = %v, want %v", got, tt.want)
			}
		})
	}
}
