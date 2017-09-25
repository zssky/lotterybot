package gen

import (
	"strconv"
	"testing"

	"github.com/zssky/lotterybot/validate"
)

func TestRed(t *testing.T) {
	t.Logf("red ball: %v", Red(2017109))
}

func TestBlue(t *testing.T) {
	t.Logf("blue ball: %v", Blue(2017109))
}

func TestWinning(t *testing.T) {
	testDBFile := "../data/db"
	v, err := validate.NewValidator(testDBFile)
	if err != nil {
		t.Fatalf("NewValidator error:%v", err)
	}

	date := 2017109
	r := Red(date)
	r2 := Red2(date)
	b := Blue(date)
	t.Logf("red ball: %v, red ball2: %v, blue ball: %v", r, r2, b)

	vr, err := v.Combinations(strconv.Itoa(date), r2, []int{}, b)
	if err != nil {
		t.Fatalf("Validate error: %v", err.Error())
	}

	t.Logf("history: %v", vr.History)
	for _, e := range vr.Entrys {
		t.Logf("match: %#v, money: %v", e.Match, e.Money)
	}
}
