package array

import "testing"

func Test_carPooling(t *testing.T) {
	type args struct {
		trips    [][]int
		capacity int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test nil",
			args: args{
				trips:    nil,
				capacity: 2,
			},
			want: true,
		},
		{
			name: "test1",
			args: args{
				trips:    [][]int{{2, 1, 5}, {3, 3, 7}},
				capacity: 4,
			},
			want: false,
		},
		{
			name: "test2",
			args: args{
				trips:    [][]int{{2, 1, 5}, {3, 3, 7}},
				capacity: 5,
			},
			want: true,
		},
		{
			name: "test2",
			args: args{
				trips:    [][]int{{2, 1, 7}, {3, 5, 7}},
				capacity: 3,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := carPooling(tt.args.trips, tt.args.capacity); got != tt.want {
				t.Errorf("carPooling() = %v, want %v", got, tt.want)
			}
		})
	}
}
