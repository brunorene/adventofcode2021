package main

import (
	"testing"
)

func Test_completeReduce(t *testing.T) {
	type args struct {
		pair string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test 1",
			args{"[[[[[9,8],1],2],3],4]"},
			"[[[[0,9],2],3],4]",
		},
		{
			"test 2",
			args{"[7,[6,[5,[4,[3,2]]]]]"},
			"[7,[6,[5,[7,0]]]]",
		},
		{
			"test 3",
			args{"[[6,[5,[4,[3,2]]]],1]"},
			"[[6,[5,[7,0]]],3]",
		},
		{
			"test 4",
			args{"[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]"},
			"[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]",
		},
		{
			"test 5",
			args{"[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]"},
			"[[3,[2,[8,0]]],[9,[5,[7,0]]]]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reduce(tt.args.pair); got != tt.want {
				t.Errorf("reduce() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_add(t *testing.T) {
	type args struct {
		pairA string
		pairB string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test 1",
			args{
				"[[[[4,3],4],4],[7,[[8,4],9]]]",
				"[1,1]",
			},
			"[[[[0,7],4],[[7,8],[6,0]]],[8,1]]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := add(tt.args.pairA, tt.args.pairB); got != tt.want {
				t.Errorf("add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_addLines(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test 1",
			args{
				[]string{
					"[1,1]",
					"[2,2]",
					"[3,3]",
					"[4,4]",
				},
			},
			"[[[[1,1],[2,2]],[3,3]],[4,4]]",
		},
		{
			"test 2",
			args{
				[]string{
					"[1,1]",
					"[2,2]",
					"[3,3]",
					"[4,4]",
					"[5,5]",
				},
			},
			"[[[[3,0],[5,3]],[4,4]],[5,5]]",
		},
		{
			"test 3",
			args{
				[]string{
					"[1,1]",
					"[2,2]",
					"[3,3]",
					"[4,4]",
					"[5,5]",
					"[6,6]",
				},
			},
			"[[[[5,0],[7,4]],[5,5]],[6,6]]",
		},
		{
			"test 4",
			args{
				[]string{
					"[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]",
					"[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]",
					"[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]",
					"[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]",
					"[7,[5,[[3,8],[1,4]]]]",
					"[[2,[2,2]],[8,[8,1]]]",
					"[2,9]",
					"[1,[[[9,3],9],[[9,0],[0,7]]]]",
					"[[[5,[7,4]],7],1]",
					"[[[[4,2],2],6],[8,7]]",
				},
			},
			"[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]",
		},
		{
			"test 5",
			args{
				[]string{
					"[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]",
					"[[[5,[2,8]],4],[5,[[9,9],0]]]",
					"[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]",
					"[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]",
					"[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]",
					"[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]",
					"[[[[5,4],[7,7]],8],[[8,3],8]]",
					"[[9,3],[[9,9],[6,[4,9]]]]",
					"[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]",
					"[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]",
				},
			},
			"[[[[6,6],[7,6]],[[7,7],[7,0]]],[[[7,7],[7,7]],[[7,8],[9,9]]]]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addLines(tt.args.lines); got != tt.want {
				t.Errorf("addLines() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_magnitude(t *testing.T) {
	type args struct {
		pair string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test 1",
			args{"[[1,2],[[3,4],5]]"},
			"143",
		},
		{
			"test 2",
			args{"[[[[0,7],4],[[7,8],[6,0]]],[8,1]]"},
			"1384",
		},
		{
			"test 3",
			args{"[[[[1,1],[2,2]],[3,3]],[4,4]]"},
			"445",
		},
		{
			"test 4",
			args{"[[[[3,0],[5,3]],[4,4]],[5,5]]"},
			"791",
		},
		{
			"test 5",
			args{"[[[[5,0],[7,4]],[5,5]],[6,6]]"},
			"1137",
		},
		{
			"test 6",
			args{"[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]"},
			"3488",
		},
		{
			"test 7",
			args{"[[[[6,6],[7,6]],[[7,7],[7,0]]],[[[7,7],[7,7]],[[7,8],[9,9]]]]"},
			"4140",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := magnitude(tt.args.pair); got != tt.want {
				t.Errorf("magnitude() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxMagnitude(t *testing.T) {
	type args struct {
		pairs []string
	}
	tests := []struct {
		name    string
		args    args
		wantMax int
	}{
		{
			"test 1",
			args{
				[]string{
					"[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]",
					"[[[5,[2,8]],4],[5,[[9,9],0]]]",
					"[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]",
					"[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]",
					"[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]",
					"[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]",
					"[[[[5,4],[7,7]],8],[[8,3],8]]",
					"[[9,3],[[9,9],[6,[4,9]]]]",
					"[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]",
					"[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]",
				},
			},
			3993,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMax := maxMagnitude(tt.args.pairs); gotMax != tt.wantMax {
				t.Errorf("maxMagnitude() = %v, want %v", gotMax, tt.wantMax)
			}
		})
	}
}
