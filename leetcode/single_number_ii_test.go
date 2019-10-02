// Problem:
// Given a list of integer where every element appears three time expect for
// one, find that unique one in O(1) space.
//
// Example:
// Input: []int{1, 1, 2, 4, 2, 6, 4, 2, 1, 4}
// Return: 6, because 6 appears 1 time only.
//
// Solution:
// Use 3 bitmask values to represent the number that appears 1 time, 2 times and
// 3 times.
// When the element appears for the third times, clear it of both the
// 1-time-appeared and 2-time-appeared bitmask values.
// The final result of 1-time-appeared bitmask value is the unique one.
//
// Cost:
// O(n) time, O(1) space.

package leetcode

import (
	"testing"

	"github.com/hoanhan101/algo/common"
)

func TestFindUniqueID(t *testing.T) {
	tests := []struct {
		in       []int
		expected int
	}{
		{[]int{}, 0},
		{[]int{6}, 6},
		{[]int{6, 2, 2, 2}, 6},
		{[]int{6, 2, 2, 2, 4, 4, 4, 1, 1, 1}, 6},
		{[]int{1, 1, 2, 4, 2, 1, 4, 6, 4, 2}, 6},
		{[]int{1, 1, 2, 4, 2, 6, 4, 2, 1, 4}, 6},
	}

	for _, tt := range tests {
		result := findUniqueID(tt.in)
		common.Equal(t, tt.expected, result)
	}
}

func findUniqueID(list []int) int {
	// ones, twos, threes represent the number that appears 1 time, 2 times and
	// 3 times accordingly.
	ones, twos, threes := 0, 0, 0

	for _, id := range list {
		// twos can be calculated by AND-ing ones and the current id.
		twos |= ones & id

		// XOR the current id with the previous ones to keep track of the
		// number that appears 1 time only.
		ones ^= id

		// when the current id appears for the third times, clear it of both
		// ones and twos.
		threes = ones & twos
		ones &= ^threes
		twos &= ^threes
	}

	return ones
}