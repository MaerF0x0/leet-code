package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var testCases = []struct {
		input1,
		input2 []int
		output bool
	}{
		{
			input1: []int{2, 1, 3, 0},
			input2: []int{2, 1, 4, 0},
			output: false,
		},
		{
			input1: []int{1, 2, 3, 4, 5},
			input2: []int{1, 2, 3, 4, 5},
			output: true,
		},
		{
			input1: []int{1, 2, 3, 4, 5},
			input2: []int{},
			output: false,
		},
		{
			input1: []int{1},
			input2: []int{1},
			output: true,
		},
	}

	for _, tc := range testCases {
		tree1, tree2 := buildTree(tc.input1), buildTree(tc.input2)
		actual := isSameTree(tree1, tree2)
		if actual != tc.output {
			fmt.Print("FAIL---")
		} else {
			fmt.Print("PASS---")
		}

		fmt.Printf(" actual=%v expected=%v \n", actual, tc.output)
	}
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if (p == nil && q != nil) || (p != nil && q == nil) {
		return false
	}
	return (p == nil && q == nil) ||
		((p.Val == q.Val) && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right))
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

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	diameter, _ := maxDepth(root)
	return diameter
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
