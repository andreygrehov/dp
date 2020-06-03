package lecture6

import "testing"

func Test_climbStairsKSteps(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "a simple test 1",
			args: args{n: 3, k: 2},
			want: 3,
		},
		{
			name: "a simple test 2",
			args: args{n: 5, k: 2},
			want: 8,
		},
		{
			name: "when n is 3",
			args: args{n: 3, k: 3},
			want: 4,
		},
		{
			name: "when n is 5",
			args: args{n: 5, k: 3},
			want: 13,
		},
		{
			name: "test when n is a big number",
			args: args{n: 1000000, k: 2},
			want: 2756670985995446685,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairsKSteps(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("climbStairsKSteps() = %v, want %v", got, tt.want)
			}
		})
	}
}