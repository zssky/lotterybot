package gen

import (
	"strconv"
	"strings"

	"github.com/zssky/lotterybot/db"
	"github.com/zssky/lotterybot/filter"
	"github.com/zssky/lotterybot/util"
)

func Red() []int {
	red := make([]int, 0)

	random := util.RandomSort(1, 33)
	killed := util.AverageSelector(random, 6)
	process := filter.Leach(random, killed)

	d, _ := db.NewSqlite3(db.DBPATH)
	ls, _ := d.GetAllHistory(map[string]string{"expect": "2017108"}, 0)

	previous := make([]int, 0)
	for _, s := range strings.Split(ls[0].Red, ",") {
		num, _ := strconv.Atoi(s)
		previous = append(previous, num)
	}

	pre := filter.Leach(process, previous)

	small, big := util.Split(pre, 17)
	left := util.AverageSelector(small, 3)
	right := util.AverageSelector(big, 3)

	red = append(append(red, left...), right...)

	return red
}

func Blue() []int {
	numbers := util.RandomSort(1, 16)
	return util.AverageSelector(numbers, 1)
}
