package main

import (
	"fmt"
)

func main() {
	A := []int{1, 2, 3, 2, 1}
	B := []int{3, 2, 1, 4, 7}

	fmt.Print(findLength(A, B))
}

func findLength(nums1 []int, nums2 []int) int {
	if len(nums1) < 1 || len(nums1) > 1000 {
		return 0
	}
	if len(nums2) < 1 || len(nums2) > 1000 {
		return 0
	}

	result := 0
	temp := make([][]int, len(nums1)+1)
	for i := range temp {
		temp[i] = make([]int, len(nums2)+1)
	}

	for i := len(nums1) - 1; i > -1; i-- {
		for j := len(nums2) - 1; j > -1; j-- {

			if nums1[i] < 0 || nums1[i] > 100 {
				return 0
			}
			if nums2[i] < 0 || nums2[i] > 100 {
				return 0
			}

			if nums1[i] == nums2[j] {
				temp[i][j] = temp[i+1][j+1] + 1
				if result < temp[i][j] {
					result = temp[i][j]
				}
			}
		}
	}

	return result
}
