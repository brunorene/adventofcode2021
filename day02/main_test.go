package main

import (
	"reflect"
	"testing"
)

func Test_move1(t *testing.T) {
	type args struct {
		movements []movement
	}
	tests := []struct {
		name    string
		args    args
		wantSub submarine
	}{
		{
			"test 1",
			args{
				[]movement{
					{forward, 5},
					{down, 5},
					{forward, 8},
					{up, 3},
					{down, 8},
					{forward, 2},
				},
			},
			submarine{15, 10, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSub := move1(tt.args.movements); !reflect.DeepEqual(gotSub, tt.wantSub) {
				t.Errorf("move1() = %v, want %v", gotSub, tt.wantSub)
			}
		})
	}
}

func Test_move2(t *testing.T) {
	type args struct {
		movements []movement
	}
	tests := []struct {
		name    string
		args    args
		wantSub submarine
	}{
		{
			"test 1",
			args{
				[]movement{
					{forward, 5},
					{down, 5},
					{forward, 8},
					{up, 3},
					{down, 8},
					{forward, 2},
				},
			},
			submarine{15, 60, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSub := move2(tt.args.movements); !reflect.DeepEqual(gotSub, tt.wantSub) {
				t.Errorf("move2() = %v, want %v", gotSub, tt.wantSub)
			}
		})
	}
}
