package lecture3

import "testing"

func Test_nSum(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "edge case #1",
			args: args{
				n: 0,
			},
			want: 0,
		},
		{
			name: "edge case #2",
			args: args{
				n: 1,
			},
			want: 1,
		},
		{
			name: "simple test #1",
			args: args{
				n: 5,
			},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nSum(tt.args.n); got != tt.want {
				t.Errorf("nSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
