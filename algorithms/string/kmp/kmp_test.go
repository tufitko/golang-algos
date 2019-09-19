package kmp

import "testing"

func TestKMP(t *testing.T) {
	type args struct {
		text string
		word string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				text: "abcxabcdabxabcdabcdabcy",
				word: "abcdabcy",
			},
			want: 15,
		}, {
			args: args{
				text: "",
				word: "",
			},
			want: 0,
		}, {
			args: args{
				text: "a",
				word: "",
			},
			want: 0,
		}, {
			args: args{
				text: "a",
				word: "a",
			},
			want: 0,
		}, {
			args: args{
				text: "abcbcglx",
				word: "abca",
			},
			want: -1,
		}, {
			args: args{
				text: "abcbcglx",
				word: "bcgl",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := KMP(tt.args.text, tt.args.word); got != tt.want {
				t.Errorf("KMP() = %v, want %v", got, tt.want)
			}
		})
	}
}
