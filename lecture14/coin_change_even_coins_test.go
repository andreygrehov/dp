package lecture14

import "testing"

func Test_coinChange(t *testing.T) {
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
			name: "simple test case #1",
			args: args{
				n:     4,
				coins: []int{1, 3, 5, 10},
			},
			want: 3, // (1,1,1,1), (1,3), (3,1)
		},
		{
			name: "simple test case #2",
			args: args{
				n:     6,
				coins: []int{1, 3, 5, 10},
			},
			want: 8, // (1,1,1,1,1,1), (1,1,1,3), (1,1,3,1), (1,3,1,1), (3,1,1,1), (3,3), (1,5), (5,1)
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := coinChange(tt.args.n, tt.args.coins); got != tt.want {
				t.Errorf("coinChange() = %v, want %v", got, tt.want)
			}
		})
	}
}
