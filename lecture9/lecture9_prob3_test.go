package lecture9

import "testing"

func Test_maxProfit(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "simple test",
			args: args{grid: [][]int{
				{0, 2, 2, 1},
				{3, 1, 1, 1},
				{4, 4, 2, 0},
			}},
			want: 13,
		},
		{
			name: "simple test 2",
			args: args{grid: [][]int{
				{0, 2, 2, 50},
				{3, 1, 1, 100},
				{4, 4, 2, 0},
			}},
			want: 154,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.grid); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}