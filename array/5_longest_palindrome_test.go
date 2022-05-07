package array

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
			name: "test empty",
			args: args{s: ""},
			want: "",
		},
		{
			name: "test one",
			args: args{s: "1"},
			want: "1",
		},
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
