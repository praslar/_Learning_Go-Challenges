package main

import "testing"

func Test(t *testing.T) {
	testCases := []struct {
		desc  string
		input []int
		want  int
	}{
		{
			desc:  "1",
			input: []int{1, 8, 3, 2},
			want:  6,
		},
		{
			desc:  "2",
			input: []int{2, 3, 3, 2},
			want:  3,
		},
		{
			desc:  "3",
			input: []int{6, 2, 4, 7},
			want:  0,
		},
		{
			desc:  "4",
			input: []int{0, 0, 0, 0, 0},
			want:  1,
		},
		{
			desc:  "5",
			input: []int{0, 0, 0, 0, 1},
			want:  5,
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
