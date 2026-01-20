package leetcode

func twoSum(nums []int, target int) []int {
	l := len(nums)

	mp := make(map[int]int)

	for i := 0; i < l; i++ {
		diff := target - nums[i]

		if j, ok := mp[diff]; ok {
			return []int{j, i}
		}
		mp[nums[i]] = i
	}
	return []int{-1, -1}
}
