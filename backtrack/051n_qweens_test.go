package main

import (
	"fmt"
	"testing"
)

func Test_solveNQueens(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test1",
			args: args{n: 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := solveNQueens(tt.args.n)
			fmt.Println(tt.name)
			for i, v := range got {
				fmt.Printf("%d>: \n", i)
				for _, v1 := range v {
					fmt.Printf("%s\n", v1)
				}
			}
		})
	}
}
