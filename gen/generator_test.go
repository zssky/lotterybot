package gen

import (
	"fmt"
	"sort"
	"strconv"
	"testing"

	"github.com/zssky/lotterybot/validate"
)

func TestRed(t *testing.T) {
	t.Logf("red ball: %v", Red(2017114))
}

func TestBlue(t *testing.T) {
	t.Logf("blue ball: %v", Blue(2017114))
}

func TestRed2(t *testing.T) {
	t.Logf("red ball: %v", Red2(2017114))
}

func TestRed3(t *testing.T) {
	t.Logf("red ball: %v", Red3(2017114))
}

func TestWinning(t *testing.T) {
	testDBFile := "../data/db"
	v, err := validate.NewValidator(testDBFile)
	if err != nil {
		t.Fatalf("NewValidator error:%v", err)
	}

	totalReport := map[string]int{
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

	redReport := map[string]int{
		"0": 0,
		"1": 0,
		"2": 0,
		"3": 0,
		"4": 0,
		"5": 0,
		"6": 0,
	}

	blueReport := map[string]int{
		"0": 0,
		"1": 0,
	}

	count := 100
	cnt := 0
	for date := 2017093; date < 2017114; date++{
		cnt++
		for i := 0; i < count; i++ {
			r := Red3(date)
			b := Blue(date)

			vr, err := v.Combinations(strconv.Itoa(date), r, []int{}, b)
			if err != nil {
				t.Fatalf("Validate error: %v", err.Error())
			}

			//t.Logf("history: %v", vr.History)
			max := validate.ValidateEntry{}
			for _, e := range vr.Entrys {
				//t.Logf("match: %#v, money: %v", e.Match, e.Money)
				if e.Match.RedCount >= max.Match.RedCount && e.Match.BlueCount >= max.Match.BlueCount {
					max = e
				}
			}

			key := fmt.Sprintf("%d+%d", max.Match.RedCount, max.Match.BlueCount)
			totalReport[key]++
			redReport[fmt.Sprintf("%d", max.Match.RedCount)]++
			blueReport[fmt.Sprintf("%d", max.Match.BlueCount)]++

			//t.Logf("Sum:%v", vr.Money)
			//if max.Money > 0 {
			//	t.Logf("max match:%#v, money: %v", max.Match, max.Money)
			//}
		}
	}

	keys := make([]string, len(totalReport))
	i := 0
	for k := range totalReport {
		keys[i] = k
		i++
	}

	sort.Strings(keys)
	for _, k := range keys {
		t.Logf("total match: %s, count: %v, rate: %f%%", k, totalReport[k], float32(totalReport[k]*100)/float32(count*cnt))
	}

	keys = make([]string, len(redReport))
	i = 0
	for k := range redReport {
		keys[i] = k
		i++
	}

	sort.Strings(keys)
	for _, k := range keys {
		t.Logf("red match: %s, count: %v, rate: %f%%", k, redReport[k], float32(redReport[k]*100)/float32(count*cnt))
	}

	keys = make([]string, len(blueReport))
	i = 0
	for k := range blueReport {
		keys[i] = k
		i++
	}

	sort.Strings(keys)
	for _, k := range keys {
		t.Logf("blue match: %s, count: %v, rate: %f%%", k, blueReport[k], float32(blueReport[k]*100)/float32(count*cnt))
	}
}
