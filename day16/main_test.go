package main

import "testing"

func Test_toBinary(t *testing.T) {
	type args struct {
		hex string
	}
	tests := []struct {
		name    string
		args    args
		wantBin string
	}{
		{
			"test 1",
			args{
				"F",
			},
			"1111",
		},
		{
			"test 1",
			args{
				"F7",
			},
			"11110111",
		},
		{
			"test 1",
			args{
				"F73",
			},
			"111101110011",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotBin := toBinary(tt.args.hex); gotBin != tt.wantBin {
				t.Errorf("toBinary() = %v, want %v", gotBin, tt.wantBin)
			}
		})
	}
}
