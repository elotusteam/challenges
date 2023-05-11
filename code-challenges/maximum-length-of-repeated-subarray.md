# Maximum Length of Repeated Subarray

Given two integer arrays **nums1** and **nums2**, return the maximum length of a subarray that appears in **both** arrays.

# Example 1:

Input: nums1 = [1,2,3,2,1], nums2 = [3,2,1,4,7]
Output: 3
Explanation: The repeated subarray with maximum length is [3,2,1].

# Example 2:

Input: nums1 = [0,0,0,0,0], nums2 = [0,0,0,0,0]
Output: 5

# Constraints:

  - **1 <= nums1.length, nums2.length <= 1000**

  - **0 <= nums1[i], nums2[i] <= 100**

# Go Template

```go
package maxsubarray

const primeRK = 16777619

func findLength(A []int, B []int) int {
	low, high := 0, min(len(A), len(B))
	for low < high {
		mid := (low + high + 1) >> 1
		if hasRepeated(A, B, mid) {
			low = mid
		} else {
			high = mid - 1
		}
	}
	return low
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func hashSlice(arr []int, length int) []int {

	hash, pl, h := make([]int, len(arr)-length+1), 1, 0
	for i := 0; i < length-1; i++ {
		pl *= primeRK
	}
	for i, v := range arr {
		h = h*primeRK + v
		if i >= length-1 {
			hash[i-length+1] = h
			h -= pl * arr[i-length+1]
		}
	}
	return hash
}

func hasSamePrefix(A, B []int, length int) bool {
	for i := 0; i < length; i++ {
		if A[i] != B[i] {
			return false
		}
	}
	return true
}

func hasRepeated(A, B []int, length int) bool {
	hs := hashSlice(A, length)
	hashToOffset := make(map[int][]int, len(hs))
	for i, h := range hs {
		hashToOffset[h] = append(hashToOffset[h], i)
	}
	for i, h := range hashSlice(B, length) {
		if offsets, ok := hashToOffset[h]; ok {
			for _, offset := range offsets {
				if hasSamePrefix(A[offset:], B[i:], length) {
					return true
				}
			}
		}
	}
	return false
}


func findLength1(A []int, B []int) int {
	res, dp := 0, make([][]int, len(A)+1)
	for i := range dp {
		dp[i] = make([]int, len(B)+1)
	}
	for i := len(A) - 1; i >= 0; i-- {
		for j := len(B) - 1; j >= 0; j-- {
			if A[i] == B[j] {
				dp[i][j] = dp[i+1][j+1] + 1
				if dp[i][j] > res {
					res = dp[i][j]
				}
			}
		}
	}
	return res
}
```

# Java Template

```java
class Solution {
    public int findLength(int[] nums1, int[] nums2) {
        
    }
}
```

# C-Sharp Template

```c#
public class Solution {
    public int FindLength(int[] nums1, int[] nums2) {
        
    }
}
```

# C++ Template

```c++
class Solution {
public:
    int findLength(vector<int>& nums1, vector<int>& nums2) {
        
    }
};
```

# Javascript Template

```js
/**
 * @param {number[]} nums1
 * @param {number[]} nums2
 * @return {number}
 */
var findLength = function(nums1, nums2) {
    
};
```
