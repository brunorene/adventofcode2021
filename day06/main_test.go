package main

import "testing"

func Test_fishCount(t *testing.T) {
	type args struct {
		line string
		days int
	}
	tests := []struct {
		name      string
		args      args
		wantCount int64
	}{
		{
			"test 1",
			args{
				"3,4,3,1,2",
				18,
			},
			26,
		},
		{
			"test 1",
			args{
				"3,4,3,1,2",
				80,
			},
			5934,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCount := fishCount(tt.args.line, tt.args.days); gotCount != tt.wantCount {
				t.Errorf("fishCount() = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}
