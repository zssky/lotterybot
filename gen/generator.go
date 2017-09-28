package gen

import (
	"fmt"
	"sort"
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
	b := util.AverageSelector(big, 2)

	red = append(append(red, s...), b...)

	left, right := util.Split(other, 17)
	l := util.AverageSelector(left, 2)
	r := util.AverageSelector(right, 1)

	red = append(append(red, l...), r...)

	sort.Ints(red)

	return red
}

func Red3(date int) []int {
	red := make([]int, 0)

	d, _ := db.NewSqlite3(db.DBPATH)
	data, _ := d.GetRedList(fmt.Sprintf(" expect<'%s' ", strconv.Itoa(date)), 2)

	count := 6
	last := make([][]int, count)
	for _, n := range data {
		for i := range n {
			last[i] = make([]int, 0)
			//if n[i]-2 > 0 {
			//	last[i] = util.AppendNum(last[i], n[i]-2)
			//}
			if n[i]-1 > 0 {
				last[i] = util.AppendNum(last[i], n[i]-1)
			}
			last[i] = util.AppendNum(last[i], n[i])
			if n[i]+1 <= 33 {
				last[i] = util.AppendNum(last[i], n[i]+1)
			}
			//if n[i]+2 <= 33 {
			//	last[i] = util.AppendNum(last[i], n[i]+2)
			//}
		}
	}

	other := util.RandomSort(1, 33)
	for _, l := range last {
		other = filter.Leach(other, l)
	}

	total := 6
	remaining := total
	otherCount := 1

	for _, l := range last {
		if remaining > otherCount {
			chance := util.AverageSelector([]int{0, 1, 2}, 1)
			if chance[0] == 0 {
				continue
			}
			b := util.AverageSelector(l, chance[0])
			for _, n := range b {
				red = util.AppendNum(red, n)
			}
			remaining = total - len(red)
		}
	}

	if remaining > 0 {
		b := util.AverageSelector(other, remaining)
		for _, n := range b {
			red = util.AppendNum(red, n)
		}
	}

	sort.Ints(red)

	return red
}

func Blue(date int) []int {
	blue := make([]int, 0)

	random := util.RandomSort(1, 16)

	d, _ := db.NewSqlite3(db.DBPATH)
	l, _ := d.GetBlueList(fmt.Sprintf(" expect<'%s' ", strconv.Itoa(date)), 4)

	killed := make([]int, 0)
	killed = append(killed, l...)

	process := filter.Leach(random, killed)
	n := util.AverageSelector(process, 1)

	blue = append(blue, n...)

	return blue
}
