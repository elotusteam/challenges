# Sum of Distances in Tree

There is an undirected connected tree with n nodes labeled from 0 to n - 1 and n - 1 edges.

You are given the integer n and the array edges where edges[i] = [ai, bi] indicates that there is an edge between nodes ai and bi in the tree.

Return an array answer of length n where answer[i] is the sum of the distances between the ith node in the tree and all other nodes.

# Example 1:

![lc-sumdist](/images/lc-sumdist.jpg)

Input: n = 6, edges = [[0,1],[0,2],[2,3],[2,4],[2,5]]
Output: [8,12,6,10,10,10]
Explanation: The tree is shown above.
We can see that dist(0,1) + dist(0,2) + dist(0,3) + dist(0,4) + dist(0,5)
equals 1 + 1 + 2 + 2 + 2 = 8.
Hence, answer[0] = 8, and so on.

# Example 2:

Input: n = 1, edges = []
Output: [0]

# Example 3:

Input: n = 2, edges = [[1,0]]
Output: [1,1]

# Constraints:

- 1 <= n <= 3 * 104
- edges.length == n - 1
- edges[i].length == 2
- 0 <= ai, bi < n
- ai != bi
- The given input represents a valid tree.

# Go Template

```go
package leetcode

func sumOfDistancesInTree(N int, edges [][]int) []int {

	tree, visited, count, res := make([][]int, N), make([]bool, N), make([]int, N), make([]int, N)
	for _, e := range edges {
		i, j := e[0], e[1]
		tree[i] = append(tree[i], j)
		tree[j] = append(tree[j], i)
	}
	deepFirstSearch(0, visited, count, res, tree)
 DFS
	visited = make([]bool, N)


	deepSecondSearch(0, visited, count, res, tree)

	return res
}

func deepFirstSearch(root int, visited []bool, count, res []int, tree [][]int) {
	visited[root] = true
	for _, n := range tree[root] {
		if visited[n] {
			continue
		}
		deepFirstSearch(n, visited, count, res, tree)
		count[root] += count[n]
)
 count[n]
		res[root] += res[n] + count[n]
	}
	count[root]++
}


func deepSecondSearch(root int, visited []bool, count, res []int, tree [][]int) {
	N := len(visited)
	visited[root] = true
	for _, n := range tree[root] {
		if visited[n] {
			continue
		}


 count[n]
 = N - count[n]

		res[n] = res[root] + (N - count[n]) - count[n]
		deepSecondSearch(n, visited, count, res, tree)
	}
}
```

# Java Template

```java
class Solution {
    public int[] sumOfDistancesInTree(int n, int[][] edges) {
        
    }
}
```

# C-Sharp Template

```c#
public class Solution {
    public int[] SumOfDistancesInTree(int n, int[][] edges) {
        
    }
}
```

# C++ Template

```c++
class Solution {
public:
    vector<int> sumOfDistancesInTree(int n, vector<vector<int>>& edges) {
        
    }
};
```

# Javascript Template

```js
/**
 * @param {number} n
 * @param {number[][]} edges
 * @return {number[]}
 */
var sumOfDistancesInTree = function(n, edges) {
    
};
```
