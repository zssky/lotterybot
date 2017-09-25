package filter

import (
	"github.com/zssky/lotterybot/util"
)

func Leach(numbers, killed []int) []int {
	f := numbers
	for _, n := range killed {
		for i := range numbers {
			if numbers[i] == n {
				f = util.Remove(f, i)
				break
			}
		}
	}

	return f
}
