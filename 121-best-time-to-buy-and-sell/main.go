package main

import (
	"fmt"
	"math"
)

type tCase struct {
	in       []int
	expected int
}

func main() {
	tCases := []tCase{
		{
			in:       []int{7, 1, 5, 3, 6, 4},
			expected: 5,
		},
		{ // cannot without enough days
			in:       []int{},
			expected: 0,
		},
		{ // cannot without enough days
			in:       []int{1},
			expected: 0,
		},
		{ // cannot earn in down market
			in:       []int{7, 6, 4, 3, 1},
			expected: 0,
		},
		{ // lower lows
			in:       []int{100, 90, 95, 80, 90, 50, 80, 30}, // 50->80 max profit
			expected: 30,
		},
	}

	for _, tcase := range tCases {
		actual := maxProfit(tcase.in)

		if actual != tcase.expected {
			fmt.Printf("FAILED -- actual=%v, expected=%v; in=%v\n", actual, tcase.expected, tcase.in)
		} else {
			fmt.Printf("PASSED -- actual=%v, expected=%v; in=%v\n", actual, tcase.expected, tcase.in)
		}
	}
}

func maxProfit(prices []int) int {
	// Cannot make money without buy and sell day
	if len(prices) < 2 {
		return 0
	}

	lowest := math.MaxInt
	highest := 0
	var curP, maxP int
	for _, price := range prices {
		if price < lowest {
			lowest = price
			highest = 0
		} else if price > highest {
			highest = price
		}
		curP = highest - lowest
		if curP > maxP {
			maxP = curP
		}
	}

	return maxP
}
