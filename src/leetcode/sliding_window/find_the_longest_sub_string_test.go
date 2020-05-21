package sliding_window

import "testing"

func Test_findTheLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"0",
			args{
				"eleetminicoworoep",
			},
			13,
		},
		{
			"1",
			args{
				"leetcodeisgreat",
			},
			5,
		},
		{
			"2",
			args{
				"bcbcbc",
			},
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTheLongestSubstring1(tt.args.s); got != tt.want {
				t.Errorf("findTheLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
