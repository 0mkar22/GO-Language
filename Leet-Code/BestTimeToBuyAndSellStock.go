package leetcode

import (
	"math"
)

func maxProfit(prices []int) int {
	minPrice := math.MaxInt32
	mp := 0
	for _, p := range prices {
		if p < minPrice {
			minPrice = p
		} else if (p - minPrice) > mp {
			mp = p - minPrice
		}
	}
	return mp
}
