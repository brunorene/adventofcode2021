package main

import "testing"

func Test_calculateRates(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name        string
		args        args
		wantGamma   int64
		wantEpsilon int64
	}{
		{
			"test 1",
			args{
				[]string{
					"00100",
					"11110",
					"10110",
					"10111",
					"10101",
					"01111",
					"00111",
					"11100",
					"10000",
					"11001",
					"00010",
					"01010",
				},
			},
			22,
			9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotGamma, gotEpsilon := calculateRates(tt.args.lines)
			if gotGamma != tt.wantGamma {
				t.Errorf("calculateRates() gotGamma = %v, want %v", gotGamma, tt.wantGamma)
			}
			if gotEpsilon != tt.wantEpsilon {
				t.Errorf("calculateRates() gotEpsilon = %v, want %v", gotEpsilon, tt.wantEpsilon)
			}
		})
	}
}

func Test_filterLines(t *testing.T) {
	type args struct {
		oxygenIn []string
		co2In    []string
	}
	tests := []struct {
		name       string
		args       args
		wantOxygen int64
		wantCo2    int64
	}{
		{
			"test 1",
			args{
				[]string{
					"00100",
					"11110",
					"10110",
					"10111",
					"10101",
					"01111",
					"00111",
					"11100",
					"10000",
					"11001",
					"00010",
					"01010",
				},
				[]string{
					"00100",
					"11110",
					"10110",
					"10111",
					"10101",
					"01111",
					"00111",
					"11100",
					"10000",
					"11001",
					"00010",
					"01010",
				},
			},
			23,
			10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOxygen, gotCo2 := filterLines(tt.args.oxygenIn, tt.args.co2In)
			if gotOxygen != tt.wantOxygen {
				t.Errorf("filterLines() gotOxygen = %v, want %v", gotOxygen, tt.wantOxygen)
			}
			if gotCo2 != tt.wantCo2 {
				t.Errorf("filterLines() gotCo2 = %v, want %v", gotCo2, tt.wantCo2)
			}
		})
	}
}
