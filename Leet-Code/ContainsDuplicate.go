package leetcode

func containsDuplicate(nums []int) bool {
	mp := make(map[int]bool)
	for _, num := range nums {
		if mp[num] {
			return true
		}
		mp[num] = true
	}
	return false
}
