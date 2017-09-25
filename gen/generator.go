package gen

import (
	"strconv"
	"strings"

	"github.com/zssky/lotterybot/db"
	"github.com/zssky/lotterybot/filter"
	"github.com/zssky/lotterybot/util"
)

func Red(date int) []int {
	red := make([]int, 0)

	random := util.RandomSort(1, 33)
	killed := util.AverageSelector(random, 6)
	process := filter.Leach(random, killed)

	d, _ := db.NewSqlite3(db.DBPATH)
	ls, _ := d.GetAllHistory(map[string]string{"expect": strconv.Itoa(date-1)}, 0)

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

func Blue(date int) []int {
	blue := make([]int, 0)

	random := util.RandomSort(1, 16)

	d, _ := db.NewSqlite3(db.DBPATH)
	l1, _ := d.GetAllHistory(map[string]string{"expect": strconv.Itoa(date-1)}, 0)
	l2, _ := d.GetAllHistory(map[string]string{"expect": strconv.Itoa(date-2)}, 0)
	l3, _ := d.GetAllHistory(map[string]string{"expect": strconv.Itoa(date-3)}, 0)

	killed := make([]int, 0)
	killed = append(killed, l1[0].Blue)
	killed = append(killed, l2[0].Blue)
	killed = append(killed, l3[0].Blue)

	process := filter.Leach(random, killed)
	n := util.AverageSelector(process, 1)

	blue = append(blue, n...)

	return blue
}
