package main

import "fmt"

func main() {

	testCases := []struct {
		n          int
		edges      [][]int
		restricted []int
		output     int
	}{
		{
			n:          7,
			edges:      [][]int{{0, 1}, {1, 2}, {3, 1}, {4, 0}, {0, 5}, {5, 6}},
			restricted: []int{4, 5},
			output:     4,
		},
		{
			n:          7,
			edges:      [][]int{{0, 1}, {0, 2}, {0, 5}, {0, 4}, {3, 2}, {6, 5}},
			restricted: []int{4, 2, 1},
			output:     3,
		},
		{
			n:          5,
			edges:      [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}},
			restricted: []int{1},
			output:     1,
		},
		{
			n:          10,
			edges:      [][]int{{4, 1}, {1, 3}, {1, 5}, {0, 5}, {3, 6}, {8, 4}, {5, 7}, {6, 9}, {3, 2}},
			restricted: []int{2, 7},
			output:     8,
		},
	}

	var failed bool
	for _, tc := range testCases {
		actual := reachableNodes(tc.n, tc.edges, tc.restricted)

		if actual != tc.output {
			fmt.Println("ERROR === expected=", tc.output, "actual=", actual, "input=", tc.n, "restricted=", tc.restricted, "edges=", tc.edges)
			failed = true
		}
	}

	if !failed {
		fmt.Println("✅ Passed all cases")
	}

}

func reachableNodes(n int, edges [][]int, restricted []int) int {
	nodes := make([]Node, n)
	restrictedMap := make(map[int]struct{}, n)
	for _, r := range restricted {
		restrictedMap[r] = struct{}{}
	}

	for _, edge := range edges {
		_, r := restrictedMap[edge[0]]
		_, r2 := restrictedMap[edge[1]]
		if r || r2 {
			continue
		}

		nodes[edge[1]].edges = append(nodes[edge[1]].edges, &nodes[edge[0]])
		nodes[edge[0]].edges = append(nodes[edge[0]].edges, &nodes[edge[1]])
	}

	return countReachable(&nodes[0], nil)
}

func countReachable(cur *Node, parent *Node) int {
	count := 1 // this node

	for _, adj := range cur.edges {
		if adj != parent {
			count += countReachable(adj, cur)
		}
	}
	return count
}

type Node struct {
	edges []*Node
}
