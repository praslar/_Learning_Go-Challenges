package main

import "testing"

func Test(t *testing.T) {
	testCases := []struct {
		desc   string
		input1 string
		input2 string
		want   string
	}{
		{
			desc:   "1",
			input1: "0",
			input2: "odd",
			want:   "IMPOSSIBLE",
		},
		{
			desc:   "2",
			input1: "test",
			input2: "tent",
			want:   "CHANGE s n",
		},
		{
			desc:   "3",
			input1: "nice",
			input2: "nicer",
			want:   "ADD r",
		},
		{
			desc:   "4",
			input1: "beans",
			input2: "banes",
			want:   "MOVE e",
		},
	}
	for _, tC := range testCases {
		got := Solution(tC.input1, tC.input2)
		t.Run(tC.desc, func(t *testing.T) {
			if got != tC.want {
				t.Error("got ", got, "but want ", tC.want)
			}
		})
	}
}
