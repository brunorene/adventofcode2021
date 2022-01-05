package main

import "testing"

func Test_findStop(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test 1",
			args{
				[]string{
					"v...>>.vv>",
					".vv>>.vv..",
					">>.>v>...v",
					">>v>>.>.v.",
					"v>v.vv.v..",
					">.>>..v...",
					".vv..>.>v.",
					"v.v..>>v.v",
					"....v..v.>",
				},
			},
			58,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findStop(tt.args.lines); got != tt.want {
				t.Errorf("findStop() = %v, want %v", got, tt.want)
			}
		})
	}
}
