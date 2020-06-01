package main

import "testing"

func Test(t *testing.T) {
	testCases := []struct {
		year  int
		desc  string
		start string
		end   string
		begin string
		want  int
	}{
		{
			desc:  "1",
			year:  2014,
			start: "April",
			end:   "May",
			begin: "Wednesday",
			want:  7,
		},
		{
			desc:  "1",
			year:  2019,
			start: "April",
			end:   "April",
			begin: "Tuesday",
			want:  4,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {

		})
	}
}
