package leetcode

func productExceptSelf(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	product := 1
	for i := 0; i < n; i++ {
		res[i] = product
		product *= nums[i]
	}
	product = 1
	for i := n - 1; i >= 0; i-- {
		res[i] *= product
		product *= nums[i]
	}
	return res
}
