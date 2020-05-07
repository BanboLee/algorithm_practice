package string

import "testing"

func Test_multiply(t *testing.T) {
	type args struct {
		num1 string
		num2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"data0",
			args{"3", "2"},
			"6",
		},
		// {
		// 	"data1",
		// 	args{"498828660196", "840477629533"},
		// 	"419254329864656431168468",
		// },
		// {
		// 	"data2",
		// 	args{"140", "721"},
		// 	"100940",
		// },
		// {
		// 	"data3",
		// 	args{"6913259244", "71103343"},
		// 	"491555843274052692",
		// },
		// {
		// 	"data4",
		// 	args{"4", "6913259244"},
		// 	"27653036976",
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := multiply1(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_add(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// {
		// 	"data0",
		// 	args{"1", "1"},
		// 	"2",
		// },
		// {
		// 	"data1",
		// 	args{"1", "9"},
		// 	"10",
		// },
		// {
		// 	"data2",
		// 	args{"1496485980588", "14964859805880"},
		// 	"16461345786468",
		// },
		// {
		// 	"data3",
		// 	args{"2940", "98000"},
		// 	"100940",
		// },
		{
			"data4",
			args{"9", "99"},
			"108",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := add1(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("add() = %v, want %v, len(add())=%d, len(want)=%d", got, tt.want, len(got), len(tt.want))
			}
		})
	}
}
