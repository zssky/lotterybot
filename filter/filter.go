package filter

func Leach(numbers, killed []int) []int {
	f := make([]int, 0)
	for i := range numbers {
		flag := true
		for _, n := range killed {
			if numbers[i] == n {
				flag = false
			}
		}
		if flag {
			f = append(f, numbers[i])
		}
	}

	return f
}
