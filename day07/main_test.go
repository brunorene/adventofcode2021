package main

import "testing"

func Test_totalFuel(t *testing.T) {
	type args struct {
		line     string
		fuelFunc func(string) int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test 1",
			args{
				"16,1,2,0,4,2,7,1,2,14",
				totalFuel,
			},
			37,
		},
		{
			"test 1",
			args{
				"16,1,2,0,4,2,7,1,2,14",
				gaussFuel,
			},
			168,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.fuelFunc(tt.args.line); got != tt.want {
				t.Errorf("totalFuel() = %v, want %v", got, tt.want)
			}
		})
	}
}
