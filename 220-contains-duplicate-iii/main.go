package main

import "fmt"

type tCase struct {
	in       []int
	k        int
	t        int
	expected bool
}

func main() {
	tCases := []tCase{
		{
			in:       []int{1, 2, 3, 1},
			k:        3,
			expected: true,
		},
		{
			in:       []int{1, 0, 1, 1},
			k:        1,
			expected: true,
		},
		{
			in:       []int{1, 2, 3, 1, 2, 3},
			k:        2,
			expected: false,
		},
		{ // edgeCase k=0 ?
			in:       []int{0, 0},
			k:        0,
			expected: false,
		},
		{ // edgeCase k > len(in) ?
			in:       []int{0, 1, 3, 4, 0},
			k:        10,
			expected: true,
		},
		{ // edgeCase duplicate at end
			in:       []int{0, 1, 2, 3, 4, 5, 6, 5},
			k:        2,
			expected: true,
		},
	}

	allCase := tCase{expected: false}
	for i := -109; i <= 109; i++ {
		allCase.in = append(allCase.in, i)
	}

	duplcase := tCase{
		in:       append(allCase.in, 109),
		k:        25, // they're side by side at the end
		expected: true,
	}

	tCases = append(tCases, allCase, duplcase)

	for _, tcase := range tCases {
		actual := containsNearbyDuplicate(tcase.in, tcase.k)

		if actual != tcase.expected {
			fmt.Printf("FAILED -- actual=%t, expected=%t; nums=%v k=%v\n", actual, tcase.expected, tcase.in, tcase.k)
		} else {
			fmt.Printf("PASSED -- actual=%t, expected=%t\n", actual, tcase.expected)
		}
	}
}

func containsNearbyDuplicate(nums []int, k int) bool {
	if k == 0 {
		// base case there is no distinct duplicate where i-j == 0
		return false
	}

	// Avoid overrunning the array
	if k > len(nums) {
		k = len(nums)
	}

	var i, j int

	seen := make(map[int]struct{}, k)
	for ; j < k; j++ {
		if _, ok := seen[nums[j]]; ok {
			return true
		}
		seen[nums[j]] = struct{}{}
	}
	// Ok, we've now covered all cases in the first k.
	// next we need to slide the window

	// NB: j retains it's value from previous for loop
	for ; j < len(nums); j++ {
		if _, ok := seen[nums[j]]; ok {
			return true
		}
		delete(seen, i) // remove one number from the sliding map we dont worry about duplicates, that's the point
		seen[nums[j]] = struct{}{}
	}

	return false
}
