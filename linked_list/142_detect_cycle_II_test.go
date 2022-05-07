package list

import (
	"reflect"
	"testing"
)

func Test_detectCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}
	cycleEnd := &ListNode{
		Val:  6,
		Next: nil,
	}
	cycleStart := &ListNode{
		Val: 3,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  5,
				Next: cycleEnd,
			},
		},
	}
	cycleEnd.Next = cycleStart
	cycleList := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: cycleStart,
		},
	}

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "test nil",
			args: args{head: nil},
			want: nil,
		},
		{
			name: "test no cycle",
			args: args{
				head: &ListNode{
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
			want: nil,
		},
		{
			name: "test cycle",
			args: args{head: cycleList},
			want: cycleStart,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectCycle(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("detectCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
