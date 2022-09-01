package main

import (
	"encoding/json"
	"fmt"
	"math"
)

const (
	NoNode int = math.MinInt
)

func main() {
	var testCases = []struct {
		input  []int
		output bool
	}{
		{
			input:  []int{2, 1, 3},
			output: true,
		},
		{
			input:  []int{1},
			output: true,
		},
		{
			input:  []int{},
			output: true,
		},
		{
			input:  []int{1, 2},
			output: true,
		},
		{
			input:  []int{1, NoNode, 2},
			output: true,
		},
		{
			input:  []int{1, 2, NoNode, 3},
			output: false,
		},
		{
			input:  []int{1, 2, NoNode, NoNode, 3},
			output: false,
		},
		{
			input:  []int{1, NoNode, 2, NoNode, NoNode, 3},
			output: false,
		},
		{
			input:  []int{1, NoNode, 2, NoNode, NoNode, NoNode, 3},
			output: false,
		},
		{
			input:  []int{1, 2, 2, 3, 3, NoNode, NoNode, 4, 4},
			output: false,
		},
	}
	for _, tc := range testCases {
		tree := buildTree(tc.input)
		actual := isBalanced(tree)

		if actual != tc.output {
			fmt.Print("FAIL---")
		} else {
			fmt.Print("PASS---")
		}

		fmt.Printf(" actual=%t expected=%t input=%v\n", actual, tc.output, tc.input)
	}
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	// If one side is nil, the other side cannot have a child
	if root.Left == nil && hasChild(root.Right) ||
		root.Right == nil && hasChild(root.Left) {
		return false
	}

	// Both not nil, recurse
	return isBalanced(root.Left) && isBalanced(root.Right)
}

func hasChild(root *TreeNode) bool {
	if root == nil {
		return false
	}
	return root.Left != nil || root.Right != nil
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(in []int) *TreeNode {
	if len(in) == 0 {
		return nil
	}

	var root *TreeNode
	q := []*(*TreeNode){&root}

	for i := 0; i < len(in); i++ {
		cur := q[0]
		if in[i] == NoNode {
			q = append(q[1:], nil, nil)
			continue
		}
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

func maxDepth(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}

	lDiameter, lDepth := maxDepth(root.Left)
	rDiameter, rDepth := maxDepth(root.Right)

	subDiameter := MaxInt(lDiameter, rDiameter)

	return MaxInt(subDiameter, lDepth+rDepth), MaxInt(lDepth, rDepth) + 1
}

func MaxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}
