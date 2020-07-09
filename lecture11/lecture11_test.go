package lecture11

import "testing"

func Test_fib(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "base case #1",
			args: args{
				n: 0,
			},
			want: 0,
		},
		{
			name: "base case #2",
			args: args{
				n: 1,
			},
			want: 1,
		},
		{
			name: "base case #3",
			args: args{
				n: 2,
			},
			want: 1,
		},
		{
			name: "simple test case",
			args: args{
				n: 10,
			},
			want: 55,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fib(tt.args.n); got != tt.want {
				t.Errorf("fib() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fibBottomUpDPBackward(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "base case #1",
			args: args{
				n: 0,
			},
			want: 0,
		},
		{
			name: "base case #2",
			args: args{
				n: 1,
			},
			want: 1,
		},
		{
			name: "base case #3",
			args: args{
				n: 2,
			},
			want: 1,
		},
		{
			name: "simple test case",
			args: args{
				n: 10,
			},
			want: 55,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fibBottomUpDPBackward(tt.args.n); got != tt.want {
				t.Errorf("fibBottomUpDPBackward() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fibBottomUpDPForward(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "base case #1",
			args: args{
				n: 0,
			},
			want: 0,
		},
		{
			name: "base case #2",
			args: args{
				n: 1,
			},
			want: 1,
		},
		{
			name: "base case #3",
			args: args{
				n: 2,
			},
			want: 1,
		},
		{
			name: "simple test case",
			args: args{
				n: 10,
			},
			want: 55,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fibBottomUpDPForward(tt.args.n); got != tt.want {
				t.Errorf("fibBottomUpDPForward() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fibTopDown(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "base case #1",
			args: args{
				n: 0,
			},
			want: 0,
		},
		{
			name: "base case #2",
			args: args{
				n: 1,
			},
			want: 1,
		},
		{
			name: "base case #3",
			args: args{
				n: 2,
			},
			want: 1,
		},
		{
			name: "simple test case",
			args: args{
				n: 10,
			},
			want: 55,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fibTopDown(tt.args.n); got != tt.want {
				t.Errorf("fibTopDown() = %v, want %v", got, tt.want)
			}
		})
	}
}