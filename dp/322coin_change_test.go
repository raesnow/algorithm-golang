package main

import "testing"

func Test_coinChange(t *testing.T) {
	type args struct {
		coins  []int
		amount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				coins:  []int{1, 2, 5},
				amount: 11,
			},
			want: 3,
		},
		{
			name: "test2",
			args: args{
				coins:  []int{2},
				amount: 3,
			},
			want: -1,
		},
		{
			name: "test3",
			args: args{
				coins:  []int{186, 419, 83, 408},
				amount: 6249,
			},
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//if got := coinChange(tt.args.coins, tt.args.amount); got != tt.want {
			//	t.Errorf("coinChange() = %v, want %v", got, tt.want)
			//}
			//if got := coinChange1(tt.args.coins, tt.args.amount); got != tt.want {
			//	t.Errorf("coinChange1() = %v, want %v", got, tt.want)
			//}
			if got := coinChange2(tt.args.coins, tt.args.amount); got != tt.want {
				t.Errorf("coinChange2() = %v, want %v", got, tt.want)
			}
		})
	}
}
