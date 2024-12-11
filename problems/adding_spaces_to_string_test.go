package problems

import (
	"testing"
)

func Test_Add_Spaces(t *testing.T) {
	testCases := [] struct {
		s    string
		spaces []int
		expected string
	} {
		{"LeetcodeHelpsMeLearn", []int{8, 13, 15}, "Leetcode Helps Me Learn"},
		{"icodeinpython", []int{1, 5, 7, 9}, "i code in py thon"},
		{"spacing", []int{0, 1, 2, 3, 4, 5, 6}, " s p a c i n g"},
	}

	for _, tc := range testCases {
		actual := AddSpaces(tc.s, tc.spaces)
		if actual != tc.expected {
			t.Errorf("Source:%s, %d\n   Expected:%s\n   Actual:  %s\n",
				tc.s,
				tc.spaces,
				tc.expected,
				actual)
		}
	}
}