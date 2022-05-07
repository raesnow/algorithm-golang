package tree

import (
	"reflect"
	"testing"
)

func Test_constructMaximumBinaryTree(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "test nil",
			args: args{nums: nil},
			want: nil,
		},
		{
			name: "test1",
			args: args{nums: []int{3, 2, 1, 6, 0, 5}},
			want: &TreeNode{
				Val: 6,
				Left: &TreeNode{
					Val:  3,
					Left: nil,
					Right: &TreeNode{
						Val:  2,
						Left: nil,
						Right: &TreeNode{
							Val:   1,
							Left:  nil,
							Right: nil,
						},
					},
				},
				Right: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val:   0,
						Left:  nil,
						Right: nil,
					},
					Right: nil,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := constructMaximumBinaryTree(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("constructMaximumBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
