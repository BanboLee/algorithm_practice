package sliding_window

import "testing"

func Test_equalSubstring(t *testing.T) {
	type args struct {
		s       string
		t       string
		maxCost int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"data1",
			args{"abcd", "bcdf", 3},
			3,
		},
		{"data2",
			args{"abcd", "cdef", 3},
			1,
		},
		{
			"data3",
			args{"abcd", "acde", 0},
			1,
		},
		{
			"data0",
			args{"ujteygggjwxnfl", "nstsenrzttikoy", 43},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := equalSubstring1(tt.args.s, tt.args.t, tt.args.maxCost); got != tt.want {
				t.Errorf("equalSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
