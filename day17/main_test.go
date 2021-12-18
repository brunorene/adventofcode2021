package main

import (
	"testing"
)

func Test_highestY(t *testing.T) {
	type args struct {
		minX int
		maxX int
		minY int
		maxY int
	}
	tests := []struct {
		name    string
		args    args
		wantMax int
	}{
		{
			"test 1",
			args{20, 30, -10, -5},
			45,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMax := highestY(tt.args.minX, tt.args.maxX, tt.args.minY, tt.args.maxY); gotMax != tt.wantMax {
				t.Errorf("highestY() = %v, want %v", gotMax, tt.wantMax)
			}
		})
	}
}

func Test_lenDistinctInitialVelocity(t *testing.T) {
	type args struct {
		minX int
		maxX int
		minY int
		maxY int
	}
	tests := []struct {
		name      string
		args      args
		wantCount int
	}{
		{
			"test 1",
			args{20, 30, -10, -5},
			112,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCount := lenDistinctInitialVelocity(tt.args.minX, tt.args.maxX, tt.args.minY, tt.args.maxY); gotCount != tt.wantCount {
				t.Errorf("lenDistinctInitialVelocity() = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}
