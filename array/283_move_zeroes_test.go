package array

import (
	"reflect"
	"testing"
)

func Test_moveZeroes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test nil",
			args: args{
				nums: nil,
			},
			want: nil,
		},
		{
			name: "test one zero",
			args: args{
				nums: []int{0},
			},
			want: []int{0},
		},
		{
			name: "test1",
			args: args{
				nums: []int{0, 1, 0, 3, 12},
			},
			want: []int{1, 3, 12, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := moveZeroes(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeDuplicates() = %v, want %v", tt.args.nums, tt.want)
			}
		})
	}
}
