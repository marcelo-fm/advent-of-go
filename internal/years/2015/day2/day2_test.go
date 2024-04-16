package main

import "testing"

func TestPartOne(t *testing.T) {
	testcases := []struct {
		in   []byte
		want int
	}{
		{[]byte("2x3x4"), 58},
		{[]byte("1x1x10"), 43},
	}
	for i, tc := range testcases {
		result, err := partOne(tc.in)
		if err != nil {
			t.Fatal(err)
		}
		if *result != tc.want {
			t.Errorf("PartOne, case %d expected %d, got %d", i+1, tc.want, result)
		}
	}
}
