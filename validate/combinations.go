package validate

func combinations(iterable []int, r int) [][]int {
	var result [][]int

	n := len(iterable)
	if r > n {
		return nil
	}

	indices := make([]int, r)
	for i := range indices {
		indices[i] = i
	}

	line := make([]int, r)
	for i, el := range indices {
		line[i] = iterable[el]
	}

	result = append(result, append([]int{}, line...))

	for i := r - 1; ; i = r - 1 {
		for ; i >= 0 && indices[i] == i+n-r; i -= 1 {
		}
		if i < 0 {
			return result
		}

		indices[i] += 1
		for j := i + 1; j < r; j += 1 {
			indices[j] = indices[j-1] + 1
		}

		for ; i < len(indices); i += 1 {
			line[i] = iterable[indices[i]]
		}
		result = append(result, append([]int{}, line...))
	}
	return result
}

//Combinations 生成组合LotteryEntry.
func Combinations(redPrefix []int, redPosfix []int, blue []int) []LotteryEntry {
	result := []LotteryEntry{}
	//需要生成组合的个数
	cs := 6 - len(redPosfix)
	css := combinations(redPrefix, cs)

	for _, line := range css {
		line = append(line, redPosfix...)
		for _, b := range blue {
			le := LotteryEntry{Blue: b}
			for i := range le.Red {
				le.Red[i] = line[i]
			}
			result = append(result, le)
		}
	}

	return result
}
