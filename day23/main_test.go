package main

import "testing"

// #############
// #...........#
// ###B#C#B#D###
//   #A#D#C#A#
//   #########

func Test_dijkstra(t *testing.T) {
	type args struct {
		start *amphiPods
	}
	tests := []struct {
		name         string
		args         args
		wantDistance int
	}{
		{
			"test 1",
			args{
				&amphiPods{typesPerPosition: map[int]int{12: 10, 11: 20, 22: 40, 21: 30, 32: 30, 31: 20, 42: 10, 41: 40}},
			},
			12521,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDistance := dijkstra(tt.args.start); gotDistance != tt.wantDistance {
				t.Errorf("dijkstra() = %v, want %v", gotDistance, tt.wantDistance)
			}
		})
	}
}
