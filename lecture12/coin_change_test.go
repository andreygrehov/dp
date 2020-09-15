package lecture12

import "testing"

func Test_coinChange(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "base case",
			args: args{n: 0},
			want: 1,
		},
		{
			name: "simple test #1",
			args: args{n: 3},
			want: 2,
		},
		{
			name: "simple test #2",
			args: args{n: 4},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := coinChange(tt.args.n); got != tt.want {
				t.Errorf("coinChange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_coinChangeWithDenomations(t *testing.T) {
	type args struct {
		n     int
		coins []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "base case",
			args: args{n: 0, coins: []int{1, 3, 5, 10}},
			want: 1,
		},
		{
			name: "simple test #1",
			args: args{n: 3, coins: []int{1, 3, 5, 10}},
			want: 2,
		},
		{
			name: "simple test #2",
			args: args{n: 4, coins: []int{1, 3, 5, 10}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := coinChangeWithDenomations(tt.args.n, tt.args.coins); got != tt.want {
				t.Errorf("coinChangeWithDenomations() = %v, want %v", got, tt.want)
			}
		})
	}
}