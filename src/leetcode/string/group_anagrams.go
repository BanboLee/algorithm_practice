package string

import "sort"

/*
字母异位词分组
给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。

示例:
输入: ["eat", "tea", "tan", "ate", "nat", "bat"],
输出:
[
  ["ate","eat","tea"],
  ["nat","tan"],
  ["bat"]
]

说明：
所有输入均为小写字母。
不考虑答案输出的顺序。

分析：
1.排序字母相等O(NKlogK) K 最长字符串长度
2.使用hash，key表示各个字母出现次数O(NK)
3. 暴力比对O(NKM) K 最长字符串长度 M 第一个字符串长度
*/
func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return nil
	}

	var res = [][]string{{strs[0]}}
	for _, str := range strs[1:] {
		isAna := false
		for i, ss := range res {
			if isAnagram(str, ss[0]) {
				ss = append(ss, str)
				res[i] = ss
				isAna = true
				break
			}
		}
		if !isAna {
			res = append(res, []string{str})
		}
	}

	return res
}

// 排序法
func groupAnagrams1(strs []string) [][]string {
	if len(strs) == 0 {
		return nil
	}
	var m = make(map[string][]string)
	for _, str := range strs {
		bs := []byte(str)
		sort.Slice(bs, func(i, j int) bool {
			return bs[i] < bs[j]
		})
		if v, ok := m[string(bs)]; ok {
			v = append(v, str)
			m[string(bs)] = v
		} else {
			m[string(bs)] = []string{str}
		}
	}
	var res [][]string
	for _, v := range m {
		res = append(res, v)
	}
	return res
}

// hash法
// hash每个字母出现的次数，相同的字符串hash一致
func groupAnagrams2(strs []string) [][]string {
	if len(strs) == 0 {
		return nil
	}

	var m = make(map[string][]string)
	for _, str := range strs {
		k := make([]byte, 26, 26)
		for _, ch := range str {
			k[ch-'a'] = k[ch-'a'] + 1
		}
		m[string(k)] = append(m[string(k)], str)
	}

	var res [][]string
	for _, v := range m {
		res = append(res, v)
	}
	return res
}
