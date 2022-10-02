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
func grayCode(n int) []int {
    
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
    // power of 2
    for (var i = 0; i < (1 << n); i++)
    {
        // Generating the decimal
        // values of gray code then using
        // bitset to convert them to binary form
        var val = (i ^ (i >> 1));
         
        // Converting to binary string
        s = val.toString(2);
        process.stdout.write(s.padStart(4, '0') + " ");
    }
};
// Driver Code
let n = 4;
   
// Function call
grayCode(n);
```
