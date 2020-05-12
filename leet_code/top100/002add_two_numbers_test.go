package main

import (
	"reflect"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	test1L1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	test1L2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	test1L3 := &ListNode{
		Val: 7,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val:  8,
				Next: nil,
			},
		},
	}

	test2L1 := &ListNode{
		Val:  5,
		Next: nil,
	}
	test2L2 := &ListNode{
		Val:  5,
		Next: nil,
	}
	test2L3 := &ListNode{
		Val: 0,
		Next: &ListNode{
			Val:  1,
			Next: nil,
		},
	}

	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "test1",
			args: args{
				l1: test1L1,
				l2: test1L2,
			},
			want: test1L3,
		},
		{
			name: "test2",
			args: args{
				l1: test2L1,
				l2: test2L2,
			},
			want: test2L3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_addTwoNumbersReverse(t *testing.T) {
	test1L1 := &ListNode{
		Val: 3,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  2,
				Next: nil,
			},
		},
	}
	test1L2 := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  5,
				Next: nil,
			},
		},
	}
	test1L3 := &ListNode{
		Val: 8,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val:  7,
				Next: nil,
			},
		},
	}

	test2L1 := &ListNode{
		Val:  5,
		Next: nil,
	}
	test2L2 := &ListNode{
		Val:  5,
		Next: nil,
	}
	test2L3 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  0,
			Next: nil,
		},
	}

	test3L1 := &ListNode{
		Val:  0,
		Next: nil,
	}
	test3L2 := &ListNode{
		Val:  5,
		Next: nil,
	}
	test3L3 := &ListNode{
		Val:  5,
		Next: nil,
	}

	test4L1 := &ListNode{
		Val:  5,
		Next: nil,
	}
	var test4L2 *ListNode = nil
	test4L3 := &ListNode{
		Val:  5,
		Next: nil,
	}

	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "test1",
			args: args{
				l1: test1L1,
				l2: test1L2,
			},
			want: test1L3,
		},
		{
			name: "test2",
			args: args{
				l1: test2L1,
				l2: test2L2,
			},
			want: test2L3,
		},
		{
			name: "test3",
			args: args{
				l1: test3L1,
				l2: test3L2,
			},
			want: test3L3,
		},
		{
			name: "test4",
			args: args{
				l1: test4L1,
				l2: test4L2,
			},
			want: test4L3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbersReverse(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbersReverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
