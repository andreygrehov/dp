package lecture10

import "testing"

func Test_numWays(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test case #1",
			args: args{
				n: 3,
			},
			want: 6,
		},
		{
			name: "test case #2",
			args: args{
				n: 4,
			},
			want: 10,
		},
		{
			name: "test case #3",
			args: args{
				n: 5,
			},
			want: 16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numWays(tt.args.n); got != tt.want {
				t.Errorf("numWays() = %v, want %v", got, tt.want)
			}
		})
	}
}