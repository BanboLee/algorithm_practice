package string

import "fmt"

/*
字符串的排列
给定两个字符串 s1 和 s2，写一个函数来判断 s2 是否包含 s1 的排列。
换句话说，第一个字符串的排列之一是第二个字符串的子串。

示例1:
输入: s1 = "ab" s2 = "eidbaooo"
输出: True
解释: s2 包含 s1 的排列之一 ("ba").

示例2:
输入: s1= "ab" s2 = "eidboaoo"
输出: False

注意：
输入的字符串只包含小写字母
两个字符串的长度都在 [1, 10,000] 之间

分析：hash + 滑动窗口判断
错误方法：所有排列、排序超时
	1. 使用hash记录s1所有可能的组合 (这里可以使用交换+递归实现，而不是每次重新排列)
	2. 使用s1大小的滑动窗口查找
正确方法：
一：记录每个单词次数为hash，依次比对每个窗口中的hash
二: 优化一，每个窗口中的hash，只需更新窗口的变化
三：优化二代码，使用数组计数
*/

// Wrong: 需要时间太长
func deepFindSeq(cur, remain []byte, n int, table map[string]uint8) {
	if len(remain) == 1 {
		cur[n] = remain[0]
		table[string(cur)] = 1
		return
	}

	var toSearch = make([][]byte, len(remain))
	for i, c := range remain {
		for j, ts := range toSearch {
			if i != j {
				ts = append(ts, c)
				toSearch[j] = ts
			}
		}
	}
	for i, c := range remain {
		cur[n] = c
		deepFindSeq(cur, toSearch[i], n+1, table)
	}
}

func checkInclusion(s1 string, s2 string) bool {
	// 1. get hash
	var table = make(map[string]uint8)
	var seq = make([]byte, len(s1))
	deepFindSeq(seq, []byte(s1), 0, table)
	fmt.Printf("%v\n", table)
	// 2. check
	w := len(s1)
	for i := 0; i <= len(s2)-w; i++ {
		println(s2[i : i+w])
		if table[s2[i:i+w]] == 1 {
			return true
		}
	}

	return false
}

// 对次数hash
func checkInclusion1(s1 string, s2 string) bool {
	var hash = make([]byte, 26, 26)
	for _, c := range s1 {
		hash[c-'a'] += 1
	}
	hs := string(hash)
	w := len(s1)
	for i := 0; i <= len(s2)-w; i++ {
		var hh = make([]byte, 26, 26)
		for _, c := range s2[i : i+w] {
			hh[c-'a'] += 1
		}
		if hs == string(hh) {
			return true
		}
	}
	return false
}

// 优化上一个
func checkInclusion2(s1 string, s2 string) bool {
	if len(s2) < len(s1) {
		return false
	}
	var hash = make([]uint8, 26, 26)
	for _, c := range s1 {
		hash[c-'a'] += 1
	}
	hs := string(hash)
	w := len(s1)
	var hh = make([]uint8, 26, 26)
	for _, ch := range s2[:w] {
		hh[ch-'a'] += 1
	}
	if hs == string(hh) {
		return true
	}
	for i := w; i < len(s2); i++ {
		hh[s2[i-w]-'a'] -= 1
		hh[s2[i]-'a'] += 1
		if hs == string(hh) {
			return true
		}
	}
	return false
}

// 优化二
func checkInclusion3(s1 string, s2 string) bool {
	var table = make(map[byte]int)
	for i := 0; i < len(s1); i++ {
		table[s1[i]]++
	}

	l, r := 0, 0
	for r < len(s2) {
		ch := s2[r]
		table[ch]-- // in
		r++
		for l < r && table[ch] < 0 { // out
			table[s2[l]]++
			l++
		}
		if r-l == len(s1) {
			return true
		}
	}
	return false
}
