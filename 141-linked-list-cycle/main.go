package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	input := buildLL([]int{1, 2, 3, 4, 5, 6})

	// third := input.Next.Next
	// last := input.Next.Next.Next.Next.Next
	// last.Next = third

	fmt.Println("input is", input)

	output := hasCycle(input)
	fmt.Println("output is", output)
}

var _ fmt.Stringer = (*ListNode)(nil)

type ListNode struct {
	Val  int
	Next *ListNode
}

var visited = make(map[*ListNode]struct{})

func (ln *ListNode) String() string {
	if _, ok := visited[ln]; ok {
		log.Println("Cycle detected and avoided infinite loop")
		return strconv.Itoa(ln.Val) // will print the node that it cycles to, but not recurse
	}

	visited[ln] = struct{}{}

	if ln == nil {
		return "nil/unknown"
	}

	if ln.Next == nil {
		return strconv.Itoa(ln.Val)
	}
	return fmt.Sprintf("%d->%s", ln.Val, ln.Next)
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	var visited = make(map[*ListNode]struct{})

	for ; head != nil; head = head.Next {
		if _, ok := visited[head]; ok {
			// We've already visited this, therefore cycle!
			return true
		}
		visited[head] = struct{}{}
	}
	return false
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
