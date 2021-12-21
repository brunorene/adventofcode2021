package main

import "testing"

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
			args{3, 7},
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
