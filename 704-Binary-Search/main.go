package main

import "fmt"

func main() {
	/*
	   Input: nums = [-1,0,3,5,9,12], target = 9
	   Output: 4
	   Explanation: 9 exists in nums and its index is 4
	   Example 2:

	   Input: nums = [-1,0,3,5,9,12], target = 2
	   Output: -1
	*/

	testCases := []struct {
		input  []int
		target int
		output int
	}{
		{ // Simple Mid case
			input:  []int{1, 2, 3},
			target: 2,
			output: 1,
		},
		{ // Simple Right case
			input:  []int{1, 2, 3},
			target: 3,
			output: 2,
		},
		{ // Simple Left case
			input:  []int{1, 2, 3},
			target: 1,
			output: 0,
		},
		{ // Left not found case
			input:  []int{1, 2, 3},
			target: -4,
			output: -1,
		},
		{ // Right not found case
			input:  []int{1, 2, 3},
			target: 400,
			output: -1,
		},
		{ // Right of mid  even case
			input:  []int{1, 2, 3, 4},
			target: 2,
			output: 1,
		},
		{ // Left of mid  even case
			input:  []int{1, 2, 3, 4},
			target: 3,
			output: 2,
		},
		{ // Less than even case
			input:  []int{1, 2, 3, 4},
			target: -400,
			output: -1,
		},
		{ // More than even case
			input:  []int{1, 2, 3, 4},
			target: 400,
			output: -1,
		},
		{ // found upper edge
			input:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 400},
			target: 400,
			output: 9,
		},
		{ // found bottom edge
			input:  []int{-400, 1, 2, 3, 4, 5, 6, 7, 8, 9, 400},
			target: -400,
			output: 0,
		},
		{
			input:  []int{-1, 0, 3, 5, 9, 12},
			target: 9,
			output: 4,
		},
		{
			input:  []int{-1, 0, 3, 5, 9, 12},
			target: 2,
			output: -1,
		},
		// Edge cases
		{
			input:  []int{}, // empty will result in not found
			output: -1,
		},
		{
			input:  []int{5}, // finds single item
			target: 5,
			output: 0,
		},
		{
			input:  []int{2, 5, 6, 7}, // finds left of mid  on even
			target: 5,
			output: 1,
		},
		{
			input:  []int{1, 2, 3, 5, 6, 7}, // finds right of mid on even
			target: 5,
			output: 3,
		},
	}

	for i, tc := range testCases {
		actual := search(tc.input, tc.target)
		if actual != tc.output {
			fmt.Printf("FAILED -- case=%d input=%#v target=%d expected=%d actual=%d\n", i, tc.input, tc.target, tc.output, actual)
		} else {
			fmt.Printf("PASS -- case=%d input=%#v target=%d expected=%d actual=%d\n", i, tc.input, tc.target, tc.output, actual)

		}
	}
}

const (
	NotFound = -1
)

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return NotFound
	}

	lower, upper := 0, len(nums)-1
	for {
		mid := ((upper - lower) / 2) + lower
		// fmt.Println("nums=", nums, "target=", target)
		// fmt.Printf("For(lower=%d mid=%d upper=%d\n", lower, mid, upper)
		if mid == upper && lower == mid && nums[mid] != target {
			return NotFound
		}

		if nums[mid] == target {
			return mid
		}

		if target < nums[mid] {
			if mid == 0 {
				return NotFound
			}
			upper = mid - 1
			continue
		}

		if target > nums[mid] {
			if mid == len(nums)-1 {
				return NotFound
			}
			lower = mid + 1
			continue
		}
	}
}
