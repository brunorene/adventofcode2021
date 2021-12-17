package main

import "testing"

func Test_sumVersions(t *testing.T) {
	type args struct {
		hex string
	}
	tests := []struct {
		name    string
		args    args
		wantSum int
	}{
		{
			"test 1",
			args{"8A004A801A8002F478"},
			16,
		},
		{
			"test 2",
			args{"620080001611562C8802118E34"},
			12,
		},
		{
			"test 3",
			args{"C0015000016115A2E0802F182340"},
			23,
		},
		{
			"test 3",
			args{"A0016C880162017C3686B18A3D4780"},
			31,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSum := sumVersions(tt.args.hex); gotSum != tt.wantSum {
				t.Errorf("sumVersions() = %v, want %v", gotSum, tt.wantSum)
			}
		})
	}
}
