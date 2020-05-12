package main

import "testing"

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{s: "babad"},
			want: "bab",
		},
		{
			name: "test2",
			args: args{s: "cbbd"},
			want: "bb",
		},
		{
			name: "test3",
			args: args{s: "a"},
			want: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestPalindrome1(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{s: "babad"},
			want: "bab",
		},
		{
			name: "test2",
			args: args{s: "cbbd"},
			want: "bb",
		},
		{
			name: "test3",
			args: args{s: "a"},
			want: "a",
		},
		{
			name: "test4",
			args: args{s: "ccc"},
			want: "ccc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome1(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestPalindrome2(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{s: "babad"},
			want: "bab",
		},
		{
			name: "test2",
			args: args{s: "cbbd"},
			want: "bb",
		},
		{
			name: "test3",
			args: args{s: "a"},
			want: "a",
		},
		{
			name: "test4",
			args: args{s: "ccc"},
			want: "ccc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome2(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome2() = %v, want %v", got, tt.want)
			}
		})
	}
}
