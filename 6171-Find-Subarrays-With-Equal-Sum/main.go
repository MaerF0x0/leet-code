package main

import "fmt"

func main() {

	tcs := []struct {
		input  []int
		output bool
	}{

		{
			input:  []int{4, 2, 4},
			output: true,
		},
		{
			input:  []int{1, 2, 3, 4, 5},
			output: false,
		},
		{
			input:  []int{0, 0, 0},
			output: true, // can be overlapping
		},
		{
			input:  []int{0, 2},
			output: false, // Cannot have distinct subarrays
		},
		{
			input:  []int{1, 2, 3, 4, 0, 5, 6, -1, 7, 8},
			output: true, // Can have negative numbers
		},
	}

	for _, tc := range tcs {
		actual := findSubarrays(tc.input)
		if actual != tc.output {
			fmt.Print("FAIL --- ")
		} else {
			fmt.Print("PASS --- ")
		}

		fmt.Println(tc.input)
	}
}

func findSubarrays(nums []int) bool {
	if len(nums) < 3 {
		return false
	}

	// Memoize by using a map of all seen sums

	// key is sum, value is the start indexes of previously seen
	seen := make(map[int]int)

	for i := 0; i < len(nums)-1; i++ {
		sum := nums[i] + nums[i+1]
		if _, ok := seen[sum]; ok {
			return true // i, idx are the two start positions
		}
		seen[sum] = i
	}

	return false

}
