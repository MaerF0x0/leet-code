package main

import "fmt"

type tCase struct {
	in       []int
	expected bool
}

func main() {
	tCases := []tCase{
		{
			in:       []int{1, 2, 3, 1},
			expected: true,
		},
		{
			in: []int{1, 2, 3, 4},

			expected: false,
		},
		{
			in:       []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			expected: true,
		},
		{ //
			in:       []int{0},
			expected: false,
		},
	}

	allCase := tCase{expected: false}
	for i := -109; i <= 109; i++ {
		allCase.in = append(allCase.in, i)
	}

	badCase := tCase{
		in:       append(allCase.in, 5),
		expected: true,
	}

	tCases = append(tCases, allCase, badCase)

	for _, tcase := range tCases {
		actual := containsDuplicate(tcase.in)

		if actual != tcase.expected {
			fmt.Printf("FAILED -- actual=%t, expected=%t\n", actual, tcase.expected)
		} else {
			fmt.Printf("PASSED -- actual=%t, expected=%t\n", actual, tcase.expected)
		}
	}
}

func containsDuplicate(nums []int) bool {
	a := make(map[int]struct{}, len(nums))

	for _, n := range nums {
		if _, ok := a[n]; ok {
			return true
		}
		a[n] = struct{}{}
	}
	return false
}
