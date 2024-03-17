# Gray Code 

An **n-bit gray** code sequence is a sequence of 2^n integers where:

  * Every integer is in the inclusive range [0, 2^n - 1],
  * The first integer is 0,
  * An integer appears no more than once in the sequence,
  * The binary representation of every pair of adjacent integers differs by exactly one bit, and
  * The binary representation of the first and last integers differs by exactly one bit.

Given an integer n, return any valid **n-bit gray** code sequence.


# Example 1:

Input: n = 2
Output: [0,1,3,2]
Explanation:
The binary representation of [0,1,3,2] is [00,01,11,10].
- 00 and 01 differ by one bit
- 01 and 11 differ by one bit
- 11 and 10 differ by one bit
- 10 and 00 differ by one bit
[0,2,3,1] is also a valid gray code sequence, whose binary representation is [00,10,11,01].
- 00 and 10 differ by one bit
- 10 and 11 differ by one bit
- 11 and 01 differ by one bit
- 01 and 00 differ by one bit

# Example 2:

Input: n = 1
Output: [0,1]

# Constraints:

1 <= n <= 16

# Go Template

```go
package graycode


func grayCode(n int) []int {
	if n == 0 {
		return []int{0}
	}
	res := []int{}
	num := make([]int, n)
	generateGrayCode(int(1<<uint(n)), 0, &num, &res)
	return res
}

func generateGrayCode(n, step int, num *[]int, res *[]int) {
	if n == 0 {
		return
	}
	*res = append(*res, convertBinary(*num))

	if step%2 == 0 {
		(*num)[len(*num)-1] = flipGrayCode((*num)[len(*num)-1])
	} else {
		index := len(*num) - 1
		for ; index >= 0; index-- {
			if (*num)[index] == 1 {
				break
			}
		}
		if index == 0 {
			(*num)[len(*num)-1] = flipGrayCode((*num)[len(*num)-1])
		} else {
			(*num)[index-1] = flipGrayCode((*num)[index-1])
		}
	}
	generateGrayCode(n-1, step+1, num, res)
	return
}

func convertBinary(num []int) int {
	res, rad := 0, 1
	for i := len(num) - 1; i >= 0; i-- {
		res += num[i] * rad
		rad *= 2
	}
	return res
}

func flipGrayCode(num int) int {
	if num == 0 {
		return 1
	}
	return 0
}


func grayCode1(n int) []int {
	var l uint = 1 << uint(n)
	out := make([]int, l)
	for i := uint(0); i < l; i++ {
		out[i] = int((i >> 1) ^ i)
	}
	return out
}
```

# Java Template

```java
class Solution {
    public List<Integer> grayCode(int n) {
        
    }
}
```
# C-Sharp Template

```c#
public class Solution {
    public IList<int> GrayCode(int n) {
        
    }
}
```

# C++ Template

```c++
class Solution {
public:
    vector<int> grayCode(int n) {
        
    }
};
```

# Javascript Template

```js
/**
 * @param {number} n
 * @return {number[]}
 */
var grayCode = function(n) {
    
};
```
