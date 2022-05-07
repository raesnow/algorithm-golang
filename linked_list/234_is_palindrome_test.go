package list

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test nil",
			args: args{head: nil},
			want: true,
		},
		{
			name: "test one node",
			args: args{head: &ListNode{
				Val:  1,
				Next: nil,
			}},
			want: true,
		},
		{
			name: "test two nodes",
			args: args{head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  2,
					Next: nil,
				},
			}},
			want: true,
		},
		{
			name: "test true even",
			args: args{head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val:  1,
							Next: nil,
						},
					},
				},
			}},
			want: true,
		},
		{
			name: "test true odd",
			args: args{head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  1,
						Next: nil,
					},
				},
			}},
			want: true,
		},
		{
			name: "test false",
			args: args{head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  3,
						Next: nil,
					},
				},
			}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.head); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
			if got := isPalindrome2(tt.args.head); got != tt.want {
				t.Errorf("isPalindrome2() = %v, want %v", got, tt.want)
			}
		})
	}
}
