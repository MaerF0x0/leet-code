package main

import "fmt"

type tCase struct {
	in       []int
	expected int
}

func main() {
	tCases := []tCase{
		{
			in:       []int{1, 7, 3, 6, 5, 6},
			expected: 3,
		},
		{
			in:       []int{1, 2, 3},
			expected: -1,
		},
		{
			in:       []int{2, 1, -1},
			expected: 0,
		},
		{ // Edge case all 0s should be leftmost index (0)
			in:       []int{0},
			expected: 0,
		},
		{ // Edge case all 0s should be leftmost index (0)
			in:       []int{0, 0},
			expected: 0,
		},
		{ // Edge case should work with the negative example too
			in:       []int{-1, -7, -3, -6, -5, -6},
			expected: 3,
		},
		{ // Ensure a pocket of negatives doesnt mess it up
			in:       []int{10, 2, 3, 30, 5, 6, 1, 1, 1, 1, -1, -1, 2}, // total = 60
			expected: 3,
			// (i=0, left=0, remaining=60, sumOfRight=50, left=10, remaining=50)
			// (i=1, sumOfRight=48, left=12, remaining=48)
			// (i=2, sumOfRight=45, left=15, remaining=45)
			// (i=3, sumOfRight=15, left=15, return 3)
		},
		{ // Sum is 0 but not at 0th index
			in:       []int{-2, 2, 0, -3, 3},
			expected: 2,
		},
		{ // Right Edge
			in:       []int{-2, 2, 0},
			expected: 2,
		},
	}

	for _, tcase := range tCases {
		actual := pivotIndex(tcase.in)

		if actual != tcase.expected {
			fmt.Printf("FAILED -- actual=%d, expected=%d\n", actual, tcase.expected)
		} else {
			fmt.Printf("PASSED -- actual=%d, expected=%d\n", actual, tcase.expected)
		}
	}
}

func pivotIndex(nums []int) int {
	// scan the whole list to make a total sum
	var total int

	for i := 0; i < len(nums); i++ {
		total += nums[i]
	}

	// Scan the list again, but find where the sum of the left, equals the right
	// (when accounting for the missing item on the pivot)
	var left int
	for i := 0; i < len(nums); i++ {
		if left == total-nums[i] {
			return i
		}
		left += nums[i]
		total -= nums[i]
	}

	return -1
}
