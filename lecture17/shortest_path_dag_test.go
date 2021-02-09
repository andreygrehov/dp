package lecture17

import "testing"

func Test_shortestPath(t *testing.T) {
	type args struct {
		graph [][]int
	}

	var graph [][]int = make([][]int, 10)
	for i := range graph {
		graph[i] = make([]int, 10)
	}

	graph[0][1] = 1
	graph[0][2] = 2
	graph[0][3] = 3

	graph[1][4] = 2
	graph[1][5] = 1
	graph[1][6] = 1

	graph[2][4] = 1
	graph[2][5] = 2
	graph[2][6] = 1

	graph[3][4] = 2
	graph[3][5] = 3
	graph[3][6] = 2

	graph[4][7] = 3
	graph[4][8] = 2

	graph[5][7] = 3
	graph[5][8] = 3

	graph[6][7] = 1
	graph[6][8] = 3

	graph[7][9] = 4
	graph[8][9] = 3

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "basic test",
			args: args{graph: graph},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestPath(tt.args.graph); got != tt.want {
				t.Errorf("shortestPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
