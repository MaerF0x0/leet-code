package main

import "fmt"

type testCase struct {
	in       []int
	expected []int
}

func main() {

	cases := []testCase{
		{
			in:       []int{1, 2, 3, 4},
			expected: []int{1, 3, 6, 10},
		},
		{
			in:       []int{1, 1, 1, 1, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			in:       []int{3, 1, 2, 10, 1},
			expected: []int{3, 4, 6, 16, 17},
		},
	}

	for i, tcase := range cases {
		actual := runningSum(tcase.in)
		if !assertEqual(tcase.in, actual) {
			fmt.Printf("FAILED - failed cases[%d]\n", i)
		} else {
			fmt.Printf("PASS - passed cases[%d]\n", i)
		}
	}
}

func runningSum(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	return nums
}

// assertEqual ensures 2 slices match
func assertEqual(expected, actual []int) bool {
	if len(actual) != len(expected) {
		fmt.Printf("FAILED - Length of actual=%d does not match expected=%d\n", len(actual), len(expected))
		return false
	}

	var success bool = true
	for idx, expVal := range expected {
		actVal := actual[idx]
		if expVal != actVal {
			fmt.Printf("FAILED - Expected[%d]=%d  != Actual[%d]=%d \n", idx, expVal, idx, actual[idx])
			success = false
		}
	}

	return success
}
