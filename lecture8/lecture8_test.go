package lecture8

import (
	"reflect"
	"testing"
)

func Test_paidStaircase(t *testing.T) {
	type args struct {
		n int
		p []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			// 0   3   2   4   6   1   1   5   3
			// |---|---|---|---|---|---|---|---|
			// 0   1   2   3   4   5   6   7   8
			//
			// 0 -> 2 -> 3 -> 5 -> 6 -> 8 = 11
			name: "complex case",
			args: args{
				n: 8,
				p: []int{0, 3, 2, 4, 6, 1, 1, 5, 3},
			},
			want: []int{0, 2, 3, 5, 6, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := paidStaircase(tt.args.n, tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("paidStaircase() = %v, want %v", got, tt.want)
			}
		})
	}
}