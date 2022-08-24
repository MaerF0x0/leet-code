package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	testCases := []struct {
		input  []int
		output []int
	}{
		{
			input:  []int{4, 2, 7, 1, 3, 6, 9},
			output: []int{4, 7, 2, 9, 6, 3, 1},
		},
		{ //
			input:  []int{2, 1, 3},
			output: []int{2, 3, 1},
		},
		{ //
			input:  []int{},
			output: []int{},
		},
		{ // single item
			input:  []int{5},
			output: []int{5},
		},
		{ // imbalanced
			input:  []int{1, 3, 4, 5},
			output: []int{1, 4, 3, 5},
		},
	}

	for _, tc := range testCases {
		reversed := invertTree(buildTree(tc.input))
		if !equalSlices(reversed.ToSliceInt(), tc.output) {
			fmt.Println("FAIL --- reversed=", reversed.ToSliceInt(), "output=", tc.output)
		} else {
			fmt.Println("PASS --- reversed=", reversed.ToSliceInt(), "output=", tc.output)
		}
	}

}

func equalSlices(l, r []int) bool {
	if len(l) != len(r) {
		return false
	}

	for i := 0; i < len(l); i++ {
		if l[i] != r[i] {
			return false
		}
	}

	return true
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Right, root.Left = invertTree(root.Left), invertTree(root.Right)
	return root
}

func buildTree(in []int) *TreeNode {
	if len(in) == 0 {
		return nil
	}

	var root *TreeNode
	q := []*(*TreeNode){&root}

	for i := 0; i < len(in); i++ {
		cur := q[0]
		*cur = &TreeNode{Val: in[i]}
		q = append(q[1:], &(*cur).Left, &(*cur).Right)
	}

	return root
}

func (tn *TreeNode) ToSliceInt() []int {
	if tn == nil {
		return []int{}
	}

	var retVal []int
	q := []*TreeNode{tn}
	var cur *TreeNode
	for len(q) > 0 {
		cur, q = q[0], q[1:]
		retVal = append(retVal, cur.Val)
		if cur.Left != nil {
			q = append(q, cur.Left)
		}
		if cur.Right != nil {
			q = append(q, cur.Right)
		}
	}

	return retVal

}

func (tn *TreeNode) String() string {
	data, err := json.MarshalIndent(tn, "", "  ")
	if err != nil {
		panic(err.Error())
	}
	return string(data)
}
