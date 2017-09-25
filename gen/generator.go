package gen

import (
	"github.com/zssky/lotterybot/util"
	"github.com/zssky/lotterybot/filter"
)

func Red() []int {
	red := make([]int, 0)

	random := util.RandomSort(1, 33)
	numbers := make([]int, len(random))
	copy(numbers, random)

	killed := util.AverageSelector(numbers, 6)
	process := filter.Leach(random, killed)

	small, big := util.Split(process, 17)
	left := util.AverageSelector(small, 3)
	right := util.AverageSelector(big, 3)

	red = append(append(red, left...), right...)

	return red
}

func Blue() []int {
	numbers := util.RandomSort(1, 16)
	return util.AverageSelector(numbers, 1)
}
