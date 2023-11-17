package main

import (
	"fmt"
	"sort"
)

func reorganizeString(s string) string {
	frequency := make(map[rune]int)

	for _, chr := range s {
		frequency[chr]++
	}

	var sortedChrs []rune
	for ch := range frequency {
		sortedChrs = append(sortedChrs, ch)
	}

	sort.Slice(sortedChrs, func(i, j int) bool {
		return frequency[sortedChrs[i]] > frequency[sortedChrs[j]]
	})

	if frequency[sortedChrs[0]] > (len(s)+1)/2 {
		return ""
	}

	result := make([]rune, len(s))
	i := 0
	for _, ch := range sortedChrs {
		for j := 0; j < frequency[ch]; j++ {
			if i >= len(s) {
				i = 1
			}
			result[i] = ch
			i += 2
		}
	}

	return string(result)
}

func main() {
	input := "aab"
	fmt.Println(reorganizeString(input))

}
