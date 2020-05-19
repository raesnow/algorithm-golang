package main

import (
	"fmt"
	"testing"
)

func Test_nQweens(t *testing.T) {
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
			got := nQweens(tt.args.n)
			fmt.Println(tt.name)
			for i, v := range got {
				fmt.Printf("%d>: \n", i)
				for _, v1 := range v {
					for _, v2 := range v1 {
						if v2 == 1 {
							fmt.Printf("* ")
						} else {
							fmt.Printf("_ ")
						}
					}
					fmt.Printf("\n")
				}
			}
		})
	}
}
