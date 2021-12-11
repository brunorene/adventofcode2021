package main

import (
	"testing"
)

func Test_doAllSteps(t *testing.T) {
	type args struct {
		lines []string
		steps int
	}
	tests := []struct {
		name      string
		args      args
		wantTotal int
	}{
		{
			"test 1",
			args{
				[]string{
					"5483143223",
					"2745854711",
					"5264556173",
					"6141336146",
					"6357385478",
					"4167524645",
					"2176841721",
					"6882881134",
					"4846848554",
					"5283751526",
				},
				100,
			},
			1656,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotTotal := doAllSteps(tt.args.lines, tt.args.steps); gotTotal != tt.wantTotal {
				t.Errorf("doAllSteps() = %v, want %v", gotTotal, tt.wantTotal)
			}
		})
	}
}

func Test_syncStep(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name          string
		args          args
		wantFlashStep int
	}{
		{
			"test 1",
			args{
				[]string{
					"5483143223",
					"2745854711",
					"5264556173",
					"6141336146",
					"6357385478",
					"4167524645",
					"2176841721",
					"6882881134",
					"4846848554",
					"5283751526",
				},
			},
			195,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotFlashStep := syncStep(tt.args.lines); gotFlashStep != tt.wantFlashStep {
				t.Errorf("syncStep() = %v, want %v", gotFlashStep, tt.wantFlashStep)
			}
		})
	}
}
