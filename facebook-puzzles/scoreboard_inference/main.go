package main

import "fmt"

func main() {

	testCases := []struct {
		name     string
		scores   []int32
		expected int32
	}{
		{
			"sample_1",
			[]int32{1, 2, 3, 4, 5, 6},
			4,
		},
		{
			"sample_2",
			[]int32{4, 3, 3, 4},
			3,
		},
		{
			"sample_3",
			[]int32{2, 4, 6, 8},
			4,
		},
		{
			"with 1",
			[]int32{1},
			1,
		},
		{
			"sample_1 reversed",
			[]int32{6, 5, 4, 3, 2, 1},
			4,
		},
		{
			"other order",
			[]int32{6, 4, 2, 3, 9}, // 2,2,2,1,2
			5,
		},
	}

	for i, tc := range testCases {
		actual := getMinProblemCount(int32(len(tc.scores)), tc.scores)

		if tc.expected != actual {
			fmt.Printf("[FAIL]:%s case[%d] expected=%d actual=%d\n", tc.name, i, tc.expected, actual)
		} else {
			fmt.Printf("[PASS]:%s case[%d]\n", tc.name, i)
		}
	}

}

func getMinProblemCount(N int32, S []int32) int32 {
	// var ones, twos int32 // how many of each problem there is

	// for _, score := range S {
	// 	if (score%2 == 1) && (ones < 1) {
	// 		// we have to add a 1 to allow odd numbers
	// 		ones++
	// 	}

	// 	// Now how much are we short?
	// 	short := score - ones - (2 * twos)

	// 	// No more left, move along
	// 	if short <= 0 {
	// 		continue
	// 	}

	// 	// we're short by > 1
	// 	twos += short / 2

	// 	// if we're still short by an odd amount
	// 	if short%2 == 1 {
	// 		ones++
	// 	}

	// 	// consolidate ones
	// 	if ones > 2 {
	// 		ones = 1
	// 		twos += ones / 2

	// 	}

	// }
	// return ones + twos

	var max, hasOdd int32 // 0 == false, 1 == true

	for _, score := range S {
		if score > max {
			max = score
		}
		if score%2 == 1 {
			hasOdd = 1
		}
	}

	return max/2 + hasOdd
}
