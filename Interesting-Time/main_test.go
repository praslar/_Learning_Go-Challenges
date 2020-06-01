package main

import "testing"

func Test(t *testing.T) {
	testCases := []struct {
		desc   string
		input1 string
		input2 string
		want   int
	}{
		{
			desc:   "1",
			input1: "15:15:00",
			input2: "15:15:12",
			want:   1,
		},
		{
			desc:   "2",
			input1: "22:22:21",
			input2: "22:22:23",
			want:   3,
		},
	}
	for _, tC := range testCases {
		got := Solution(tC.input1, tC.input2)
		t.Run(tC.desc, func(t *testing.T) {
			if got != tC.want {
				t.Error("got ", got, " but want ", tC.want)
			}
		})
	}
}
