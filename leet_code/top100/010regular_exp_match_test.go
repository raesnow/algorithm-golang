package main

import "testing"

func Test_isMatch(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{
				s: "aa",
				p: "a",
			},
			want: false,
		},
		{
			name: "test2",
			args: args{
				s: "aa",
				p: "a*",
			},
			want: true,
		},
		{
			name: "test3",
			args: args{
				s: "ab",
				p: ".*",
			},
			want: true,
		},
		{
			name: "test4",
			args: args{
				s: "aab",
				p: "c*a*b",
			},
			want: true,
		},
		{
			name: "test5",
			args: args{
				s: "mississippi",
				p: "mis*is*p*.",
			},
			want: false,
		},
		{
			name: "test6",
			args: args{
				s: "",
				p: "mis*is*p*.",
			},
			want: false,
		},
		{
			name: "test7",
			args: args{
				s: "",
				p: "",
			},
			want: true,
		},
		{
			name: "test7",
			args: args{
				s: "bbbba",
				p: ".*a*a",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
