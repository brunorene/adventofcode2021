package main

import (
	"reflect"
	"testing"
)

func Test_countDotsAfter1stFold(t *testing.T) {
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
					"6,10",
					"0,14",
					"9,10",
					"0,3",
					"10,4",
					"4,11",
					"6,0",
					"6,12",
					"4,1",
					"0,13",
					"10,12",
					"3,4",
					"3,0",
					"8,4",
					"1,10",
					"2,14",
					"8,10",
					"9,0",
					"",
					"fold along y=7",
					"fold along x=5",
				},
			},
			17,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countDotsAfter1stFold(tt.args.lines); got != tt.want {
				t.Errorf("countDotsAfter1stFold() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_capitalLetters(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test 1",
			args{
				[]string{
					"6,10",
					"0,14",
					"9,10",
					"0,3",
					"10,4",
					"4,11",
					"6,0",
					"6,12",
					"4,1",
					"0,13",
					"10,12",
					"3,4",
					"3,0",
					"8,4",
					"1,10",
					"2,14",
					"8,10",
					"9,0",
					"",
					"fold along y=7",
					"fold along x=5",
				},
			},
			`#####
#...#
#...#
#...#
#####
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := capitalLetters(tt.args.lines).String(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("capitalLetters()\n\n%s\nwant\n\n%s\n", got, tt.want)
			}
		})
	}
}
