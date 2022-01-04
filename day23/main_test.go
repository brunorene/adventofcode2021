package main

import "testing"

func Test_dijkstra(t *testing.T) {
	type args struct {
		start []string
	}
	tests := []struct {
		name         string
		args         args
		wantDistance int
	}{
		{
			"test 1",
			args{
				[]string{
					"#############",
					"#...........#",
					"###B#C#B#D###",
					"  #A#D#C#A#",
					"  #########",
				},
			},
			12521,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDistance := dijkstra(processInput(tt.args.start)); gotDistance != tt.wantDistance {
				t.Errorf("dijkstra() = %v, want %v", gotDistance, tt.wantDistance)
			}
		})
	}
}
