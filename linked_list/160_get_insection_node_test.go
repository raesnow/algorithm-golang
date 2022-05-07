package list

import (
	"reflect"
	"testing"
)

func Test_getIntersectionNode(t *testing.T) {
	type args struct {
		headA *ListNode
		headB *ListNode
	}
	commonNodes := &ListNode{
		Val: 6,
		Next: &ListNode{
			Val:  7,
			Next: nil,
		},
	}

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "test nil",
			args: args{
				headA: nil,
				headB: &ListNode{
					Val:  1,
					Next: nil,
				},
			},
			want: nil,
		},
		{
			name: "test no intersection",
			args: args{
				headA: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  8,
						Next: nil,
					},
				},
				headB: &ListNode{
					Val:  1,
					Next: nil,
				},
			},
			want: nil,
		},
		{
			name: "test has intersection 2",
			args: args{
				headA: &ListNode{
					Val:  1,
					Next: commonNodes,
				},
				headB: commonNodes,
			},
			want: commonNodes,
		},
		{
			name: "test has intersection",
			args: args{
				headA: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val:  2,
						Next: commonNodes,
					},
				},
				headB: &ListNode{
					Val:  5,
					Next: commonNodes,
				},
			},
			want: commonNodes,
		},
		{
			name: "test same list",
			args: args{
				headA: commonNodes,
				headB: commonNodes,
			},
			want: commonNodes,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getIntersectionNode(tt.args.headA, tt.args.headB); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getIntersectionNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
