package main

import (
	"fmt"
	"sort"
)

func main() {
	var tcs = []struct {
		input    [][]int
		expected int
	}{
		{
			input: [][]int{
				{5, 10},
				{6, 8},
				{1, 5},
				{2, 3},
				{1, 10},
			},
			expected: 3,
		},
		{
			input: [][]int{
				{1, 3}, {5, 6}, {8, 10}, {11, 13},
			},
			expected: 1,
		},
	}

	for i, tc := range tcs {
		actual := minGroups(tc.input)
		if actual != tc.expected {
			fmt.Print("FAIL---")
		} else {
			fmt.Print("PASS---")
		}

		fmt.Printf("case=%d, expected=%d actual=%d \n", i, tc.expected, actual)
	}
}

const (
	LeftValue  = 0
	RightValue = 1
)

type interval struct {
	subIntervals [][]int
	min          int
	max          int
}

func (i *interval) String() string {
	return fmt.Sprintf("Interval(%d->%d (%+v)", i.min, i.max, i.subIntervals)
}

func (i *interval) canAdd(check []int) bool {
	// fmt.Println("canAdd", i, check, check[LeftValue] > i.max)
	return check[LeftValue] > i.max
}

func (i *interval) add(subInt []int) bool {
	if len(i.subIntervals) == 0 {
		// fmt.Println("setting i", i, ".min=", i.min, subInt[LeftValue])
		i.min = subInt[LeftValue]
	}

	i.subIntervals = append(i.subIntervals, subInt)
	i.max = subInt[RightValue]
	return true
}

func minGroups(intervals [][]int) int {
	intvs := sortIntervals(intervals)
	sort.Sort(intvs)

	groups := []*interval{{subIntervals: [][]int{}}}

	for _, ival := range intvs {
		var added bool
		for _, group := range groups {
			if group.canAdd(ival) {
				group.add(ival)
				added = true
				break
			}
		}

		if !added {
			newGroup := &interval{}
			newGroup.add(ival)
			groups = append(groups, newGroup)
		}
	}

	// fmt.Println(groups)
	return len(groups)
}

type sortIntervals [][]int

func (si sortIntervals) Len() int {
	return len(si)
}

func (si sortIntervals) Less(i, j int) bool {
	if si[i][LeftValue] == si[j][LeftValue] {
		return si[i][RightValue] <= si[j][RightValue]
	}

	return si[i][LeftValue] < si[j][LeftValue]

}

func (si sortIntervals) Swap(i, j int) {
	si[i], si[j] = si[j], si[i]
}
