package tree

import (
	"reflect"
	"testing"
)

func Test_connect(t *testing.T) {
	type args struct {
		root *Node
	}
	rightRight := &Node{
		Val:   7,
		Left:  nil,
		Right: nil,
		Next:  nil,
	}
	rightLeft := &Node{
		Val:   6,
		Left:  nil,
		Right: nil,
		Next:  rightRight,
	}
	leftRight := &Node{
		Val:   5,
		Left:  nil,
		Right: nil,
		Next:  rightLeft,
	}
	right := &Node{
		Val:   3,
		Left:  rightLeft,
		Right: rightRight,
		Next:  nil,
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{
			name: "test nil",
			args: args{root: nil},
			want: nil,
		},
		{
			name: "test1",
			args: args{root: &Node{
				Val: 1,
				Left: &Node{
					Val: 2,
					Left: &Node{
						Val:   4,
						Left:  nil,
						Right: nil,
						Next:  nil,
					},
					Right: &Node{
						Val:   5,
						Left:  nil,
						Right: nil,
						Next:  nil,
					},
					Next: nil,
				},
				Right: &Node{
					Val: 3,
					Left: &Node{
						Val:   6,
						Left:  nil,
						Right: nil,
						Next:  nil,
					},
					Right: &Node{
						Val:   7,
						Left:  nil,
						Right: nil,
						Next:  nil,
					},
					Next: nil,
				},
				Next: nil,
			}},
			want: &Node{
				Val: 1,
				Left: &Node{
					Val: 2,
					Left: &Node{
						Val:   4,
						Left:  nil,
						Right: nil,
						Next:  leftRight,
					},
					Right: leftRight,
					Next:  right,
				},
				Right: right,
				Next:  nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := connect(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("connect() = %v, want %v", got, tt.want)
			}
		})
	}
}
