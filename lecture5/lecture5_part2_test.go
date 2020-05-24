package lecture5

import "testing"

func Test_climbStairs3Steps(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "when n is 3",
			args: args{n: 3},
			want: 4,
		},
		{
			name: "when n is 5",
			args: args{n: 5},
			want: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairs3Steps(tt.args.n); got != tt.want {
				t.Errorf("climbStairs3Steps() = %v, want %v", got, tt.want)
			}
		})
	}
}
