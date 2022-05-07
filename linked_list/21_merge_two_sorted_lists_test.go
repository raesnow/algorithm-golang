package list

import (
	"reflect"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		list1 *ListNode
		list2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "test both nil",
			args: args{
				list1: nil,
				list2: nil,
			},
			want: nil,
		},
		{
			name: "test one nil",
			args: args{
				list1: nil,
				list2: &ListNode{
					Val:  1,
					Next: nil,
				},
			},
			want: &ListNode{
				Val:  1,
				Next: nil,
			},
		},
		{
			name: "test has values",
			args: args{
				list1: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val:  3,
						Next: nil,
					},
				},
				list2: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val:  5,
							Next: nil,
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.args.list1, tt.args.list2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
