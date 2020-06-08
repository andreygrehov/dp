package lecture7

import "testing"

func Test_paidStaircase(t *testing.T) {
	type args struct {
		n int
		p []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "simple test case",
			args: args{
				n: 3,
				p: []int{0, 3, 2, 4},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := paidStaircase(tt.args.n, tt.args.p); got != tt.want {
				t.Errorf("paidStaircase() = %v, want %v", got, tt.want)
			}
		})
	}
}