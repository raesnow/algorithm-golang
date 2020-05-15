package main

import "testing"

func Test_isValidBST(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{root: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 3},
			}},
			want: true,
		},
		{
			name: "test2",
			args: args{root: &TreeNode{
				Val:  5,
				Left: &TreeNode{Val: 1},
				Right: &TreeNode{
					Val:   4,
					Left:  &TreeNode{Val: 3},
					Right: &TreeNode{Val: 6},
				},
			}},
			want: false,
		},
		{
			name: "test3",
			args: args{root: nil},
			want: true,
		},
		{
			name: "test3",
			args: args{root: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 1},
				Right: nil,
			}},
			want: false,
		},
		{
			name: "test4",
			args: args{root: &TreeNode{
				Val:   1,
				Left:  nil,
				Right: &TreeNode{Val: 1},
			}},
			want: false,
		},
		{
			name: "test4",
			args: args{root: &TreeNode{
				Val:  3,
				Left: nil,
				Right: &TreeNode{
					Val: 30,
					Left: &TreeNode{
						Val:  10,
						Left: nil,
						Right: &TreeNode{
							Val:   15,
							Left:  nil,
							Right: &TreeNode{Val: 45},
						},
					},
					Right: nil,
				},
			}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBST(tt.args.root); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
