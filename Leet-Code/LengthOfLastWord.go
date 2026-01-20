package leetcode

import (
	"strings"
)

func lengthOfLastWord(s string) int {
	slicewords := strings.Fields(s)

	return len(slicewords[len(slicewords)-1])
}
