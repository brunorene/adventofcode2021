package main

import "testing"

func Test_mostMinusless(t *testing.T) {
	type args struct {
		lines     []string
		stepCount int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			"test 1",
			args{
				[]string{
					"NNCB",
					"",
					"CH -> B",
					"HH -> N",
					"CB -> H",
					"NH -> C",
					"HB -> C",
					"HC -> B",
					"HN -> C",
					"NN -> C",
					"BH -> H",
					"NC -> B",
					"NB -> B",
					"BN -> B",
					"BB -> N",
					"BC -> B",
					"CC -> N",
					"CN -> C",
				},
				10,
			},
			int64(1588),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mostMinusless(tt.args.lines, tt.args.stepCount); got != tt.want {
				t.Errorf("mostMinusless() = %v, want %v", got, tt.want)
			}
		})
	}
}
