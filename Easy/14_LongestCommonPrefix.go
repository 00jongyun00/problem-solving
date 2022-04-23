package main

import (
	"sort"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	sort.Slice(strs, func(i, j int) bool {
		return len(strs[i]) < len(strs[j])
	})
	target := strings.Split(strs[0], "")
	result := ""

	for _, t := range target {
		current := result + t
		flag := true
		for _, s := range strs {
			if !strings.HasPrefix(s, current) {
				flag = false
			}
		}
		if flag {
			result += t
		}
	}
	return result
}
