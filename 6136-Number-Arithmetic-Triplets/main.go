package main

import "fmt"

func main() {

	testCases := []struct {
		input  []int
		diff   int
		output int
	}{
		{
			input:  []int{0, 1, 2, 4, 6, 7, 10},
			diff:   3,
			output: 2,
		},
		{
			input:  []int{4, 5, 6, 7, 8, 9},
			diff:   2,
			output: 2,
		},
		{
			// Min case
			input:  []int{1, 2, 3},
			diff:   1,
			output: 1,
		},
	}

	var failed bool
	for _, tc := range testCases {
		actual := arithmeticTriplets(tc.input, tc.diff)

		if actual != tc.output {
			fmt.Println("ERROR === expected=", tc.output, "actual=", actual, "input=", tc.input, "diff=", tc.diff)
			failed = true
		}
	}

	if !failed {
		fmt.Println("✅ Passed all cases")
	}

}

func arithmeticTriplets(nums []int, diff int) int {
	var count int
	var checks int
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				// TODO equate these?
				checks++
				// fmt.Println("checking ", i, j, k)
				if ((nums[j] - nums[i]) == diff) && ((nums[k] - nums[j]) == diff) {
					count++
				}
			}
		}
	}
	fmt.Println("checks ran ", checks)
	return count
}
