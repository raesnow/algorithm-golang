package array

import "testing"

func TestNumArray_SumRange(t *testing.T) {
	type fields struct {
		nums []int
	}
	type args struct {
		left  int
		right int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "test nil",
			fields: fields{
				nums: nil,
			},
			args: args{
				left:  0,
				right: 1,
			},
			want: 0,
		},
		{
			name: "test args invalid",
			fields: fields{
				nums: []int{1, 4, 6},
			},
			args: args{
				left:  2,
				right: 1,
			},
			want: 0,
		},
		{
			name: "test1",
			fields: fields{
				nums: []int{-2, 0, 3, -5, 2, -1},
			},
			args: args{
				left:  0,
				right: 2,
			},
			want: 1,
		},
		{
			name: "test2",
			fields: fields{
				nums: []int{-2, 0, 3, -5, 2, -1},
			},
			args: args{
				left:  2,
				right: 5,
			},
			want: -1,
		},
		{
			name: "test3",
			fields: fields{
				nums: []int{-2, 0, 3, -5, 2, -1},
			},
			args: args{
				left:  0,
				right: 5,
			},
			want: -3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := Constructor(tt.fields.nums)
			if got := this.SumRange(tt.args.left, tt.args.right); got != tt.want {
				t.Errorf("SumRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
