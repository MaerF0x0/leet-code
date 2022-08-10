package main

import "fmt"

func main() {
	testCases := []struct {
		nums   []int
		output int
	}{
		{
			nums:   []int{10, 9, 2, 5, 3, 7, 101, 18},
			output: 4,
		},
		{
			nums:   []int{0, 1, 0, 3, 2, 3},
			output: 4,
		},
		{
			nums:   []int{7, 7, 7, 7, 7, 7, 7},
			output: 1,
		},
		{
			nums:   []int{7},
			output: 1,
		},
		{
			nums:   []int{7, 6, 5, 4, 3, 1, -100},
			output: 1,
		},
		{
			nums:   []int{1, 2, 3, 4, 5},
			output: 5,
		},
		{
			nums:   []int{1, 2, 3, 4, 100, 100, 100, 1, 2, 3, 4, 5, 6},
			output: 6, // 1,2,3,4,5,6
		},
	}

	var failed bool
	for _, tc := range testCases {
		actual := lengthOfLIS(tc.nums)

		if actual != tc.output {
			fmt.Println("ERROR === expected=", tc.output, "actual=", actual, "nums=", tc.nums)
			failed = true
		}
	}

	if !failed {
		fmt.Println("✅ Passed all cases")
	}

}

func lengthOfLIS(nums []int) int {
	var sub []int

	for _, n := range nums {
		if len(sub) == 0 || sub[len(sub)-1] < n {
			sub = append(sub, n)
		} else {
			i := lowerBound(sub, n)
			sub[i] = n
		}

	}
	return len(sub)
}

func lowerBound(in []int, n int) int {
	for i, n2 := range in {
		if n2 >= n {
			return i
		}
	}
	// This should never hit
	return 0
}
