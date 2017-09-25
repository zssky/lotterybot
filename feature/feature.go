package feature

import (
	"github.com/juju/errors"
	"github.com/zssky/lotterybot/db"
)

func GetSum(nums []int) int {
	sum := 0
	for i := range nums {
		sum += nums[i]
	}

	return sum
}
