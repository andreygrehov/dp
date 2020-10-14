package lecture16

import "testing"

func Test_changeMaking(t *testing.T) {
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
				n:     29,
				coins: []int{1, 3, 5},
			},
			want: 7,
		},
		{
			name: "simple test case #2",
			args: args{
				n:     1,
				coins: []int{2, 3, 5},
			},
			want: -1,
		},
		{
			name: "simple test case #3",
			args: args{
				n:     56,
				coins: []int{15, 4, 3},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := changeMaking(tt.args.n, tt.args.coins); got != tt.want {
				t.Errorf("changeMaking() = %v, want %v", got, tt.want)
			}
		})
	}
}
