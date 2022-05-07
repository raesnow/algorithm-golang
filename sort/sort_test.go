package sort

import (
	"reflect"
	"testing"
)

func Test_sortArray(t *testing.T) {
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
			args: args{nums: nil},
			want: nil,
		},
		{
			name: "test1",
			args: args{nums: []int{5, 1, 1, 2, 0, 0}},
			want: []int{0, 0, 1, 1, 2, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortArray(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortArray() = %v, want %v", got, tt.want)
			}
			if got := sortArray1(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortArray1() = %v, want %v", got, tt.want)
			}
			if got := sortArray2(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortArray2() = %v, want %v", got, tt.want)
			}
			if got := sortArray3(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortArray3() = %v, want %v", got, tt.want)
			}
			if got := sortArray4(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortArray4() = %v, want %v", got, tt.want)
			}
		})
	}
}
