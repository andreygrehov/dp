package lecture5

import "testing"

func Test_climbStairs(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "a simple test 1",
			args: args{n: 3},
			want: 3,
		},
		{
			name: "a simple test 2",
			args: args{n: 5},
			want: 8,
		},
		{
			name: "test memory when n is a big number",
			args: args{n: 1000000},
			want: 2756670985995446685,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairs(tt.args.n); got != tt.want {
				t.Errorf("climbStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
