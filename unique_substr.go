package main

import (
	"fmt"
)

func LongestUniqueStr(a string) string {
	runes := []rune(a)

	set := make(map[rune]struct{})
	l := 0
	max_so_far := 0
	start := 0

	for r := 0; r < len(runes); r++ {
		if _, exist := set[runes[r]]; exist {
			delete(set, runes[r])
			l += 1
		}

		set[runes[r]] = struct{}{}

		if (r - l + 1) > max_so_far {
			max_so_far = r - l + 1
			start = l
		}
	}

	return string(runes[start : start+max_so_far])
}

func main() {
	s := []string{"abcabcbb", "bbbbb", "abababbab"}

	for _, valstr := range s {
		fmt.Println(LongestUniqueStr(valstr))
	}
}
