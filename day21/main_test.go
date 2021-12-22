package main

import (
	"testing"
)

func Test_determPlay(t *testing.T) {
	type args struct {
		p1Pos int
		p2Pos int
	}
	tests := []struct {
		name       string
		args       args
		wantResult int
	}{
		{
			"test 1",
			args{4, 8},
			739785,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := determPlay(tt.args.p1Pos, tt.args.p2Pos); gotResult != tt.wantResult {
				t.Errorf("determPlay() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func Test_startGame(t *testing.T) {
	type args struct {
		p1Start int
		p2Start int
	}
	tests := []struct {
		name       string
		args       args
		wantP1Wins int64
		wantP2Wins int64
	}{
		{
			"test 1",
			args{4, 8},
			444356092776315,
			341960390180808,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotP1Wins, gotP2Wins := startGame(tt.args.p1Start, tt.args.p2Start)
			if gotP1Wins != tt.wantP1Wins {
				t.Errorf("startGame() gotP1Wins = %v, want %v", gotP1Wins, tt.wantP1Wins)
			}
			if gotP2Wins != tt.wantP2Wins {
				t.Errorf("startGame() gotP2Wins = %v, want %v", gotP2Wins, tt.wantP2Wins)
			}
		})
	}
}
