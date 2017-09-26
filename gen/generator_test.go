package gen

import (
	"fmt"
	"sort"
	"strconv"
	"testing"

	"github.com/zssky/lotterybot/validate"
)

func TestRed(t *testing.T) {
	t.Logf("red ball: %v", Red(2017113))
}

func TestBlue(t *testing.T) {
	t.Logf("blue ball: %v", Blue(2017113))
}

func TestRed2(t *testing.T) {
	t.Logf("red ball: %v", Red2(2017113))
}

func TestWinning(t *testing.T) {
	testDBFile := "../data/db"
	v, err := validate.NewValidator(testDBFile)
	if err != nil {
		t.Fatalf("NewValidator error:%v", err)
	}

	count := 1000
	date := 2017109
	report := map[string]int{
		"0+0": 0,
		"1+0": 0,
		"2+0": 0,
		"3+0": 0,
		"4+0": 0,
		"5+0": 0,
		"6+0": 0,
		"0+1": 0,
		"1+1": 0,
		"2+1": 0,
		"3+1": 0,
		"4+1": 0,
		"5+1": 0,
		"6+1": 0,
	}

	sum := 0
	for i := 0; i < count; i++ {
		//r := Red(date)
		r2 := Red2(date)
		b := Blue(date)
		//t.Logf("red ball: %v, red ball2: %v, blue ball: %v", r, r2, b)

		vr, err := v.Combinations(strconv.Itoa(date), r2, []int{}, b)
		if err != nil {
			t.Fatalf("Validate error: %v", err.Error())
		}

		//t.Logf("history: %v", vr.History)
		max := validate.ValidateEntry{}
		for _, e := range vr.Entrys {
			//t.Logf("match: %#v, money: %v", e.Match, e.Money)
			if e.Money > max.Money {
				max = e
			}
			key := fmt.Sprintf("%d+%d", e.Match.RedCount, e.Match.BlueCount)
			report[key]++
			sum++
		}

		//t.Logf("Sum:%v", vr.Money)
		//if max.Money > 0 {
		//	t.Logf("max match:%#v, money: %v", max.Match, max.Money)
		//}
	}

	keys := make([]string, len(report))
	i := 0
	for k := range report {
		keys[i] = k
		i++
	}

	sort.Strings(keys)
	for _, k := range keys {
		t.Logf("match: %s, count: %v, rate: %f%%", k, report[k], float32(report[k]*100/sum))
	}
}
