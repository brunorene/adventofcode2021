package main

import (
	"reflect"
	"testing"
)

func Test_allCandidates(t *testing.T) {
	type args struct {
		min    int
		max    int
		inside func(vel, min, max int) []motionData
	}
	tests := []struct {
		name           string
		args           args
		wantCandidates map[motionData]struct{}
	}{
		{
			"test X",
			args{20, 30, positionsInsideX},
			map[motionData]struct{}{},
		},
		{
			"test Y",
			args{-10, -5, positionsInsideY},
			map[motionData]struct{}{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCandidates := allCandidates(tt.args.min, tt.args.max, tt.args.inside); !reflect.DeepEqual(gotCandidates, tt.wantCandidates) {
				t.Errorf("allCandidates() = %v, want %v", gotCandidates, tt.wantCandidates)
			}
		})
	}
}

func Test_highestY(t *testing.T) {
	type args struct {
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
			args{-10, -5},
			45,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMax := highestY(tt.args.minY, tt.args.maxY); gotMax != tt.wantMax {
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
		name string
		args args
		want int
	}{
		{
			"test 1",
			args{20, 30, -10, -5},
			112,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lenDistinctInitialVelocity(tt.args.minX, tt.args.maxX, tt.args.minY, tt.args.maxY); got != tt.want {
				t.Errorf("lenDistinctInitialVelocity() = %v, want %v", got, tt.want)
			}
		})
	}
}
