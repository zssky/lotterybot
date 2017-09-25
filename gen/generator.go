package gen

import (
	"github.com/zssky/lotterybot/util"
)

func Red() []int {
	return util.Random(6, 33)
}

func Blue() []int {
	return util.Random(1, 16)
}
