package main

import (
	"testing"
)

func Test_countIncreases(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test 1",
			args{
				[]int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263},
			},
			7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countIncreases(tt.args.numbers); got != tt.want {
				t.Errorf("countIncreases() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countIncreasesWithSlidingWindow(t *testing.T) {
	type args struct {
		size    int
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test 1",
			args{
				3,
				[]int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263},
			},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countIncreasesWithSlidingWindow(tt.args.size, tt.args.numbers); got != tt.want {
				t.Errorf("countIncreasesWithSlidingWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}
