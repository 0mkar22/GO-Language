package leetcode

import "strconv"

func isPalindrome(x int) bool {
	xstr := strconv.Itoa(x)

	for ix, jx := 0, len(xstr)-1; ix <= jx; {
		if xstr[ix] != xstr[jx] {
			return false
		}
		ix++
		jx--
	}
	return true

}
