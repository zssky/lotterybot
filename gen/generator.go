package gen

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/zssky/lotterybot/db"
	"github.com/zssky/lotterybot/filter"
	"github.com/zssky/lotterybot/util"
	"sort"
)

func Red(date int) []int {
	red := make([]int, 0)

	random := util.RandomSort(1, 33)
	killed := util.AverageSelector(random, 6)
	process := filter.Leach(random, killed)

	d, _ := db.NewSqlite3(db.DBPATH)
	ls, _ := d.GetAllHistory(map[string]string{"expect": strconv.Itoa(date - 1)}, 0)

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

func Red2(date int) []int {
	red := make([]int, 0)

	random := util.RandomSort(1, 33)
	//killed := util.AverageSelector(random, 6)
	//process := filter.Leach(random, killed)

	d, _ := db.NewSqlite3(db.DBPATH)
	data, _ := d.GetRedList(fmt.Sprintf(" expect<'%s' ", strconv.Itoa(date)), 6)

	last := make([]int, 0)
	for i := range data {
		for j := range data[i] {
			flag := true
			//for k := range killed {
			//	if killed[k] == data[i][j] {
			//		flag = false
			//		break
			//	}
			//}
			if flag {
				last = util.AppendNum(last, data[i][j])
			}
		}
	}

	other := filter.Leach(random, last)

	small, big := util.Split(last, 17)
	s := util.AverageSelector(small, 3)
	b := util.AverageSelector(big, 3)

	red = append(append(red, s...), b...)

	left, right := util.Split(other, 17)
	l := util.AverageSelector(left, 2)
	r := util.AverageSelector(right, 1)

	red = append(append(red, l...), r...)

	sort.Ints(red)

	return red
}

func Blue(date int) []int {
	blue := make([]int, 0)

	random := util.RandomSort(1, 16)

	d, _ := db.NewSqlite3(db.DBPATH)
	l, _ := d.GetBlueList(fmt.Sprintf(" expect<'%s' ", strconv.Itoa(date)), 3)

	killed := make([]int, 0)
	killed = append(killed, l...)

	process := filter.Leach(random, killed)
	n := util.AverageSelector(process, 1)

	blue = append(blue, n...)

	return blue
}
