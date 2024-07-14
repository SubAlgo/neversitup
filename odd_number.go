package main

func FindOddNumber(text []int) int {
	check := make(map[int]int)

	for _, v := range text {
		check[v] += 1
	}

	var res int
	for k, v := range check {
		if v%2 != 0 {
			res = k
			break
		}
	}
	return res
}
