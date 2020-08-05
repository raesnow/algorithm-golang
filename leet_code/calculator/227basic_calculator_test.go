package calculator

import "testing"

func Test_calculate1(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{s: "221 + (-542 + 321)"},
			want: 0,
		},
		{
			name: "test2",
			args: args{s: "1 + 1"},
			want: 2,
		},
		{
			name: "test3",
			args: args{s: " 2-1 + 2 "},
			want: 3,
		},
		{
			name: "test4",
			args: args{s: "(1+(4+5+2)-3)+(6+8)"},
			want: 23,
		},
		{
			name: "test5",
			args: args{s: "3+2*2"},
			want: 7,
		},
		{
			name: "test6",
			args: args{s: "3/2"},
			want: 1,
		},
		{
			name: "test7",
			args: args{s: "3 + 5/2"},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculate1(tt.args.s); got != tt.want {
				t.Errorf("calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
