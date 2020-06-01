package main

import "testing"

func Test(t *testing.T) {
	testCases := []struct {
		input1 []string
		input2 []string
		want   int
		desc   string
	}{
		{
			desc:   "Test1",
			input1: []string{"test1a", "test2", "test1b"},
			input2: []string{"Wrong answer", "OK", "Runtime error"},
			want:   50,
		},
		{
			desc:   "Test2",
			input1: []string{"codility1", "codility3", "codility2", "codility4b", "codility4a"},
			input2: []string{"Wrong answer", "OK", "OK", "Runtime error", "OK"},
			want:   50,
		},
	}
	for _, tC := range testCases {
		got := Solution(tC.input1, tC.input2)
		t.Run(tC.desc, func(t *testing.T) {
			if tC.want != got {
				t.Error("got ", got, " but want ", tC.want)
			}
		})
	}
}
