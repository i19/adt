package leetcode

import "testing"

func Test_minimumTotal(t *testing.T) {
	type args struct {
		triangle [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1", args: args{triangle: [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}}, want: 11},
		{name: "test2", args: args{triangle: [][]int{{-10}}}, want: -10},
		{name: "test3", args: args{triangle: [][]int{{2}, {3, 4}, {6, 5, 9}, {4, 4, 8, 0}}}, want: 14},
		{name: "test4", args: args{triangle: [][]int{{-7}, {-2, 1}, {-5, -5, 9}, {-4, -5, 4, 4}, {-6, -6, 2, -1, -5}, {3, 7, 8, -3, 7, -9}, {-9, -1, -9, 6, 9, 0, 7}, {-7, 0, -6, -8, 7, 1, -4, 9}, {-3, 2, -6, -9, -7, -6, -9, 4, 0}, {-8, -6, -3, -9, -2, -6, 7, -5, 0, 7}, {-9, -1, -2, 4, -2, 4, 4, -1, 2, -5, 5}, {1, 1, -6, 1, -2, -4, 4, -2, 6, -6, 0, 6}, {-3, -3, -6, -2, -6, -2, 7, -9, -5, -7, -5, 5, 1}}}, want: -63},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumTotal(tt.args.triangle); got != tt.want {
				t.Errorf("minimumTotal() = %v, want %v", got, tt.want)
			}
		})
	}
}
