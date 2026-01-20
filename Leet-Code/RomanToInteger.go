package leetcode

func romanToInt(s string) int {
	total := 0

	rom := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	for i, v := range s {
		total += rom[string(v)]
		if i != 0 {
			if rom[string(s[i-1])] < rom[string(v)] {
				total -= 2 * rom[string(s[i-1])]
			}
		}
	}
	return total

}
