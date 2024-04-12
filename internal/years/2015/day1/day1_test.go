package day1

import "testing"

func TestPartOne(t *testing.T) {
	testcases := []struct {
		in   string
		want int
	}{
		{"./input/test/partOne/caseOne.txt", 0},
		{"./input/test/partOne/caseTwo.txt", 3},
		{"./input/test/partOne/caseThree.txt", 3},
		{"./input/test/partOne/caseFour.txt", -1},
		{"./input/test/partOne/caseFive.txt", -3},
	}
	for _, tc := range testcases {
		result := partOne(tc.in)
		if result != tc.want {
			t.Errorf("PartOne produced: %d, want %d", result, tc.want)
		}
	}
}

func TestPartTwo(t *testing.T) {
	testcases := []struct {
		in   string
		want int
	}{
		{"./input/test/partTwo/caseOne.txt", 1},
		{"./input/test/partTwo/caseTwo.txt", 5},
	}
	for _, tc := range testcases {
		result := partTwo(tc.in)
		if result != tc.want {
			t.Errorf("PartTwo produced: %d, want %d", result, tc.want)
		}
	}
}
