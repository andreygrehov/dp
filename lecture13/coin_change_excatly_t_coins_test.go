package lecture13

import "testing"

func Test_coinChangeExactlyTCoins(t *testing.T) {
	type args struct {
		n     int
		t     int
		coins []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "base case",
			args: args{
				n:     0,
				t:     0,
				coins: []int{1, 2, 3, 5},
			},
			want: 1,
		},
		{
			name: "simple test",
			args: args{
				n:     7,
				t:     3,
				coins: []int{1, 2, 3, 5},
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := coinChangeExactlyTCoins(tt.args.n, tt.args.t, tt.args.coins); got != tt.want {
				t.Errorf("coinChangeExactlyTCoins() = %v, want %v", got, tt.want)
			}
		})
	}
}
