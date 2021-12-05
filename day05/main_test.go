package main

import "testing"

func Test_countOverlaps(t *testing.T) {
	type args struct {
		lines  []string
		filter func(int, int, int, int) bool
	}
	tests := []struct {
		name      string
		args      args
		wantCount int
	}{
		{
			"test not diagonal",
			args{
				[]string{
					"0,9 -> 5,9",
					"8,0 -> 0,8",
					"9,4 -> 3,4",
					"2,2 -> 2,1",
					"7,0 -> 7,4",
					"6,4 -> 2,0",
					"0,9 -> 2,9",
					"3,4 -> 1,4",
					"0,0 -> 8,8",
					"5,5 -> 8,2",
				},
				notDiagonal,
			},
			5,
		},
		{
			"test all",
			args{
				[]string{
					"0,9 -> 5,9",
					"8,0 -> 0,8",
					"9,4 -> 3,4",
					"2,2 -> 2,1",
					"7,0 -> 7,4",
					"6,4 -> 2,0",
					"0,9 -> 2,9",
					"3,4 -> 1,4",
					"0,0 -> 8,8",
					"5,5 -> 8,2",
				},
				anyLine,
			},
			12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCount := countOverlaps(tt.args.lines, tt.args.filter); gotCount != tt.wantCount {
				t.Errorf("countOverlaps() = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}
