package string

import (
	"math"
	"testing"
)

func Test_myAtoi(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"data0",
			args{"-91283472332"},
			math.MinInt32,
		},
		{
			"data1",
			args{"4193 with words"},
			4193,
		},
		{
			"data2",
			args{"2147483648"},
			2147483647,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := atoiDFA(tt.args.str); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
