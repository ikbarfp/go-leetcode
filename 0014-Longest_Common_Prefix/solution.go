package main

func longestCommonPrefix(strs []string) string {
	first := strs[0]
	for _, s := range strs {
		i := 0
		for ; i < len(s) && i < len(first) && first[i] == s[i]; i++ {
		}
		first = first[:i]
	}
	return first
}
