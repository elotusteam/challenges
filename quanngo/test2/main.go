package main

import (
	"fmt"
)

func main() {
	input := [][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}, {2, 5}}

	fmt.Print(sumOfDistancesInTree(6, input))
}

func sumOfDistancesInTree(n int, edges [][]int) []int {
	if n < 1 && n > 3*104 {
		return nil
	}
	if len(edges) != n-1 {
		return nil
	}

	sum := make([]int, n)
	if len(edges) != n-1 {
		return nil
	}

	count := make([]int, n)
	input := make(map[int][]int)
	for index, edge := range edges {

		if (len(edges[index])) != 2 {
			return nil
		}

		u, v := edge[0], edge[1]
		input[u] = append(input[u], v)
		input[v] = append(input[v], u)
	}
	var postOrder func(int, int)
	postOrder = func(src, parent int) {
		for _, adj := range input[src] {
			if adj == parent {
				continue
			}
			postOrder(adj, src)
			count[src] += count[adj]
			sum[src] += sum[adj] + count[adj]
		}
		count[src] += 1
	}
	postOrder(0, -1)
	var preOrder func(int, int)
	preOrder = func(src, parent int) {
		for _, adj := range input[src] {
			if parent == adj {
				continue
			}
			sum[adj] = sum[src] - count[adj] + n - count[adj]
			preOrder(adj, src)
		}
	}
	preOrder(0, -1)
	return sum
}
