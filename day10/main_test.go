package main

import (
	"testing"
)

func Test_totalScore(t *testing.T) {
	type args struct {
		lines []string
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
					"[({(<(())[]>[[{[]{<()<>>",
					"[(()[<>])]({[<{<<[]>>(",
					"{([(<{}[<>[]}>{[]{[(<()>",
					"(((({<>}<{<{<>}{[]{[]{}",
					"[[<[([]))<([[{}[[()]]]",
					"[{[{({}]{}}([{[{{{}}([]",
					"{<[[]]>}<{[{[{[]{()[[[]",
					"[<(<(<(<{}))><([]([]()",
					"<{([([[(<>()){}]>(<<{{",
					"<{([{{}}[<[[[<>{}]]]>[]]",
				},
			},
			26397,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotTotal := totalErrorScore(tt.args.lines); gotTotal != tt.wantTotal {
				t.Errorf("totalScore() = %v, want %v", gotTotal, tt.wantTotal)
			}
		})
	}
}

func Test_totalAutocompleteScore(t *testing.T) {
	type args struct {
		lines []string
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
					"[({(<(())[]>[[{[]{<()<>>",
					"[(()[<>])]({[<{<<[]>>(",
					"{([(<{}[<>[]}>{[]{[(<()>",
					"(((({<>}<{<{<>}{[]{[]{}",
					"[[<[([]))<([[{}[[()]]]",
					"[{[{({}]{}}([{[{{{}}([]",
					"{<[[]]>}<{[{[{[]{()[[[]",
					"[<(<(<(<{}))><([]([]()",
					"<{([([[(<>()){}]>(<<{{",
					"<{([{{}}[<[[[<>{}]]]>[]]",
				},
			},
			288957,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotTotal := totalAutocompleteScore(tt.args.lines); gotTotal != tt.wantTotal {
				t.Errorf("totalAutocompleteScore() = %v, want %v", gotTotal, tt.wantTotal)
			}
		})
	}
}
