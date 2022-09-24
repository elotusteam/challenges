package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Print(grayCode(2))
}

func grayCode(n int) []int {
	if 1 <= n && n <= 16 {
		bin := int(math.Pow(2, float64(n)))
		result := make([]int, bin)
		for i := range result {
			result[i] = i ^ (i >> 1)
		}
		return result
	}

	return nil
}
