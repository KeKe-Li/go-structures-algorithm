# [3.Longest Substring Without Repeating Characters](https://leetcode.com/problems/longest-substring-without-repeating-characters/)

## 题目

Given a string, find the length of the longest substring without repeating characters.



Example 1:

```c
Input: "abcabcbb"
Output: 3 
Explanation: The answer is "abc", with the length of 3. 
```

Example 2:

```c
Input: "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
```

Example 3:

```c
Input: "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3. 
             Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
```

题意:

在一个字符串重寻找没有重复字母的最长子串。

## 解题思路

利用s[left:i+1]来表示s[:i+1]中的包含s[i]的最长子字符串。location[s[i]]是字符s[i]在s[:i+1]中倒数第二次出现的序列号。
当left < location[s[i]]的时候，说明字符s[i]出现了两次。需要设置 left = location[s[i]] + 1,保证字符s[i]只出现一次。


滑动窗口的右边界不断的右移，只要没有重复的字符，就不用的向右扩大窗口边界。一旦出现了重复字符，此时先计算一下滑动窗口的大小，记录下来。再需要缩小左边界。直到重复的字符移出了左边界。接着又可以开始移动滑动窗口的右边界。以此类推，不断的刷新记录的窗口大小。最终最大的值就是题目中的所求。


## 总结
利用Location保存字符上次出现的序列号，避免了查询工作。location和[Two Sum](./algorithms/0001.two-sum)中的m是一样的作用。


### go实现
```go
func lengthOfLongestSubstring(s string) int {
	// 哈希集合，记录每个字符是否出现过
	m := map[byte]int{}
	n := len(s)
	// 右指针，初始值为 -1，相当于我们在字符串的左边界的左侧，还没有开始移动
	rk, ans := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			// 左指针向右移动一格，移除一个字符
			delete(m, s[i-1])
		}
		for rk + 1 < n && m[s[rk+1]] == 0 {
			// 不断地移动右指针
			m[s[rk+1]]++
			rk++
		}
		// 第 i 到 rk 个字符是一个极长的无重复字符子串
		ans = max(ans, rk - i + 1)
	}
	return ans
}


func max(x,y int) int{
	if x <y {
		return y
	}
	return x
}

```








