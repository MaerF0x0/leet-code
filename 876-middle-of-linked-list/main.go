package main

import (
	"fmt"
	"strconv"
)

func main() {
	input := buildLL([]int{1, 2, 3, 4, 5, 6})

	output := middleNode(input)
	fmt.Println("output is", output)
}

var _ fmt.Stringer = (*ListNode)(nil)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (ln *ListNode) String() string {
	if ln == nil {
		return "nil/unknown"
	}

	if ln.Next == nil {
		return strconv.Itoa(ln.Val)
	}
	return fmt.Sprintf("%d->%s", ln.Val, ln.Next)
}

func middleNode(head *ListNode) *ListNode {
	mid := head
	for idx := 0; head != nil; idx++ {
		// every other we advance mid
		if idx%2 == 1 {
			fmt.Printf("advancing mid %s\n", mid)
			mid = mid.Next
		}
		head = head.Next
	}

	return mid
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var prev *ListNode

	for head != nil {
		head, head.Next, prev = head.Next, prev, head
	}

	return prev
}

func buildLL(input []int) *ListNode {
	if len(input) == 0 {
		return nil
	}
	head := ListNode{
		Val: input[0],
	}
	cur := &head
	for _, val := range input[1:] {

		cur.Next = &ListNode{
			Val: val,
		}

		cur = cur.Next
	}

	return &head
}

func equalLL(a, b *ListNode) bool {
	aCur, bCur := a, b
	for {
		if aCur.Val != bCur.Val {
			return false
		}

		// Everything up to this point has been equal value wise
		if aCur.Next == nil && bCur.Next == nil {
			return true
		} else if aCur.Next == nil || bCur.Next == nil {
			// One is nil, but not both. That means the length doesnt match
			return false
		}

		// advance pointer to next node and repeat
		aCur, bCur = aCur.Next, bCur.Next
	}
}
