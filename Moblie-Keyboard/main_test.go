package main

import "testing"

func Test(t *testing.T) {
	testCases := []struct {
		desc  string
		input string
		want  string
	}{
		{
			desc:  "1",
			input: "CrCellBax",
			want:  "Relax",
		},
	}
	for _, tC := range testCases {
		got := Solution(tC.input)
		t.Run(tC.desc, func(t *testing.T) {
			if got != tC.want {
				t.Error("got", got, " but want", tC.want)
			}
		})
	}
}
