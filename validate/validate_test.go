package validate

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/juju/errors"

	"github.com/zssky/lotterybot/db"
)

var (
	testDBFile      string
	testLotteryData = []db.Lottery{
		{Expect: "2017111", Red: "05,10,17,19,29,32", Blue: 12, OpenTime: "2017-09-21 21:18:20", OpenTimestamp: "1505999900"},
		{Expect: "2017110", Red: "01,03,12,15,19,23", Blue: 14, OpenTime: "2017-09-19 21:18:20", OpenTimestamp: "1505827100"},
		{Expect: "2017109", Red: "08,14,16,18,21,23", Blue: 16, OpenTime: "2017-09-17 21:18:20", OpenTimestamp: "1505654300"},
		{Expect: "2017108", Red: "07,12,14,15,17,20", Blue: 1, OpenTime: "2017-09-14 21:18:20", OpenTimestamp: "1505395100"},
		{Expect: "2017107", Red: "08,09,15,17,30,32", Blue: 6, OpenTime: "2017-09-12 21:18:20", OpenTimestamp: "1505222300"},
		{Expect: "2017106", Red: "12,15,20,25,27,31", Blue: 2, OpenTime: "2017-09-10 21:18:20", OpenTimestamp: "1505049500"},
		{Expect: "2017105", Red: "03,06,07,12,25,26", Blue: 7, OpenTime: "2017-09-07 21:18:20", OpenTimestamp: "1504790300"},
		{Expect: "2017104", Red: "01,14,15,20,23,30", Blue: 14, OpenTime: "2017-09-05 21:18:20", OpenTimestamp: "1504617500"},
		{Expect: "2017103", Red: "01,21,23,25,31,33", Blue: 1, OpenTime: "2017-09-03 21:18:20", OpenTimestamp: "1504444700"},
		{Expect: "2017102", Red: "04,08,10,14,18,20", Blue: 11, OpenTime: "2017-08-31 21:18:20", OpenTimestamp: "1504185500"},
		{Expect: "2017101", Red: "01,04,11,28,31,32", Blue: 16, OpenTime: "2017-08-29 21:18:20", OpenTimestamp: "1504012700"},
		{Expect: "2017100", Red: "04,07,08,18,23,24", Blue: 2, OpenTime: "2017-08-27 21:18:20", OpenTimestamp: "1503839900"},
		{Expect: "2017099", Red: "02,05,06,16,28,29", Blue: 4, OpenTime: "2017-08-24 21:18:20", OpenTimestamp: "1503580700"},
		{Expect: "2017098", Red: "04,19,22,27,30,33", Blue: 1, OpenTime: "2017-08-22 21:18:20", OpenTimestamp: "1503407900"},
		{Expect: "2017097", Red: "05,10,18,19,30,31", Blue: 3, OpenTime: "2017-08-20 21:18:20", OpenTimestamp: "1503235100"},
		{Expect: "2017096", Red: "02,06,11,12,19,29", Blue: 6, OpenTime: "2017-08-17 21:18:20", OpenTimestamp: "1502975900"},
		{Expect: "2017095", Red: "09,10,12,19,22,29", Blue: 16, OpenTime: "2017-08-15 21:18:20", OpenTimestamp: "1502803100"},
		{Expect: "2017094", Red: "08,11,13,19,28,31", Blue: 6, OpenTime: "2017-08-13 21:18:20", OpenTimestamp: "1502630300"},
		{Expect: "2017093", Red: "07,08,09,15,22,27", Blue: 12, OpenTime: "2017-08-10 21:18:20", OpenTimestamp: "1502371100"},
		{Expect: "2017092", Red: "10,18,19,29,32,33", Blue: 9, OpenTime: "2017-08-08 21:18:20", OpenTimestamp: "1502198300"},
	}
)

func TestMain(main *testing.M) {
	testDBFile = fmt.Sprintf("/tmp/lotterybot_%v.db", time.Now().UnixNano())
	ldb, err := db.NewSqlite3(testDBFile)
	if err != nil {
		panic(fmt.Sprintf("NewSqlite3 %v, error:%v", testDBFile, err))
	}

	_, err = ldb.Exec("CREATE TABLE history ( `id` INTEGER PRIMARY KEY AUTOINCREMENT, `expect` TEXT, `red` TEXT, `blue`  INTEGER, `opentime` TEXT, `opentimestamp` NUMERIC )")
	if err != nil {
		panic(err.Error())
	}

	for _, ld := range testLotteryData {
		if err = ldb.AddHistory(ld); err != nil {
			panic(fmt.Sprintf("AddHistory %v error:%v", ld, err))
		}
	}

	ldb.Close()

	main.Run()

	os.Remove(testDBFile)
}

func TestValidate(t *testing.T) {
	v, err := NewValidator(testDBFile)
	if err != nil {
		t.Fatalf("NewValidator error:%v", err)
	}

	tds := []struct {
		entry LotteryEntry
		money int
	}{
		//全没中
		{LotteryEntry{Red: [6]int{1, 2, 3, 4, 5, 6}, Blue: 1}, 0},
		//六等奖 0+1,1+1,2+1 都是5块
		{LotteryEntry{Red: [6]int{1, 2, 3, 4, 5, 6}, Blue: 9}, 5},
		{LotteryEntry{Red: [6]int{1, 2, 3, 4, 5, 33}, Blue: 9}, 5},
		{LotteryEntry{Red: [6]int{1, 2, 3, 4, 32, 33}, Blue: 9}, 5},

		//五等奖 3+1,4+0 都是10块
		{LotteryEntry{Red: [6]int{1, 2, 3, 29, 32, 33}, Blue: 9}, 10},
		{LotteryEntry{Red: [6]int{1, 2, 19, 29, 32, 33}, Blue: 1}, 10},

		//四等奖 4+1,5+0 都是200块钱
		{LotteryEntry{Red: [6]int{1, 2, 19, 29, 32, 33}, Blue: 9}, 200},
		{LotteryEntry{Red: [6]int{1, 18, 19, 29, 32, 33}, Blue: 1}, 200},

		//三等奖 5+1 是3000
		{LotteryEntry{Red: [6]int{1, 18, 19, 29, 32, 33}, Blue: 9}, 3000},

		//二等奖 6+0 当期高等奖奖金的30%， 这里写死的150000
		{LotteryEntry{Red: [6]int{10, 18, 19, 29, 32, 33}, Blue: 1}, 150000},

		//一等奖 6+1 最高奖，这里写死5000000
		{LotteryEntry{Red: [6]int{10, 18, 19, 29, 32, 33}, Blue: 9}, 5000000},
	}

	for _, td := range tds {
		vr, err := v.Validate("2017092", []LotteryEntry{td.entry})
		if err != nil {
			t.Fatalf("Validate error:%v, entry:%v", errors.Trace(err), td.entry)
		}
		if td.money != vr.Money {
			t.Fatalf("expect money:%v, recv:%v, data:%#v", td.money, vr.Money, vr)
		}
	}
}

func TestValidateMultiple(t *testing.T) {
	v, err := NewValidator(testDBFile)
	if err != nil {
		t.Fatalf("NewValidator error:%v", err)
	}

	tds := []struct {
		entry []LotteryEntry
		money int
	}{
		{[]LotteryEntry{
			LotteryEntry{Red: [6]int{1, 2, 3, 4, 5, 6}, Blue: 9},
			LotteryEntry{[6]int{1, 2, 3, 29, 32, 33}, 9},
			LotteryEntry{[6]int{1, 2, 19, 29, 32, 33}, 9},
			LotteryEntry{[6]int{1, 18, 19, 29, 32, 33}, 9},
		}, 5 + 10 + 200 + 3000},
	}

	for _, td := range tds {
		vr, err := v.Validate("2017092", td.entry)
		if err != nil {
			t.Fatalf("Validate error:%v, entry:%v", errors.Trace(err), td.entry)
		}
		if td.money != vr.Money {
			t.Fatalf("expect money:%v, recv:%v, data:%#v", td.money, vr.Money, vr)
		}
		t.Logf("history:%v", vr.History)
		for _, e := range vr.Entrys {
			t.Logf("entry:%#v, match:%#v, money:%v", e.Entry, e.Match, e.Money)
		}
	}
}
