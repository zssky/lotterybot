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
	b := Blue(date)
	t.Logf("red ball: %v, blue ball: %v", r, b)
	td := []validate.LotteryEntry{
		{Red: [6]int{r[0], r[1], r[2], r[3], r[4], r[5]}, Blue: b[0]},
	}

	vr, err := v.Validate(strconv.Itoa(date), td)
	if err != nil {
		t.Fatalf("Validate error: %v, entry: %v", err.Error(), td)
	}

	t.Logf("history: %v", vr.History)
	for _, e := range vr.Entrys {
		t.Logf("match: %#v, money: %v", e.Match, e.Money)
	}
}
