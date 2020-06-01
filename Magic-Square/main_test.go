package main

import "testing"

func Test(t *testing.T) {
	testCases := []struct {
		desc  string
		input [][]int
		want  int
	}{
		{
			desc: "4x5",
			input: [][]int{
				{4, 3, 4, 5, 3},
				{2, 7, 3, 8, 4},
				{1, 7, 6, 5, 2},
				{8, 4, 9, 5, 5},
			},
			want: 3,
		},
		{
			desc: "2x4",
			input: [][]int{
				{2, 2, 1, 1},
				{2, 2, 2, 2},
				{1, 2, 2, 2},
			},
			want: 2,
		},
		{
			desc: "5x3",
			input: [][]int{
				{7, 2, 4},
				{2, 7, 6},
				{9, 5, 1},
				{4, 3, 8},
				{3, 5, 4},
			},
			want: 3,
		},
		{
			desc: "1x1",
			input: [][]int{
				{7},
			},
			want: 1,
		},
	}
	for _, tC := range testCases {
		got := Solution(tC.input)
		t.Run(tC.desc, func(t *testing.T) {
			if got != tC.want {
				t.Error("got ", got, " but want ", tC.want)
			}
		})
	}
}
