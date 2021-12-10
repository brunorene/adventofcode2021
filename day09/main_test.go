package main

import (
	"testing"
)

func Test_lowValues(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name    string
		args    args
		wantSum int
	}{
		{
			"test 1",
			args{
				[]string{
					"2199943210",
					"3987894921",
					"9856789892",
					"8767896789",
					"9899965678",
				},
			},
			15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSum := sumLowValues(tt.args.lines); gotSum != tt.wantSum {
				t.Errorf("lowValues() = %v, want %v", gotSum, tt.wantSum)
			}
		})
	}
}

func Test_basins(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name     string
		args     args
		wantMult int
	}{
		{
			"test 1",
			args{
				[]string{
					"2199943210",
					"3987894921",
					"9856789892",
					"8767896789",
					"9899965678",
				},
			},
			1134,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMult := basins(tt.args.lines); gotMult != tt.wantMult {
				t.Errorf("basins() = %v, want %v", gotMult, tt.wantMult)
			}
		})
	}
}
