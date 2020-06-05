package list

import "testing"

func Test_isPalindrome1(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"0",
			args{&ListNode{0, &ListNode{0, nil}}},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome1(tt.args.head); got != tt.want {
				t.Errorf("isPalindrome1() = %v, want %v", got, tt.want)
			}
		})
	}
}
