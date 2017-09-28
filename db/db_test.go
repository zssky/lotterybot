package db

import (
	"encoding/json"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
	"testing"
)

const (
	ssq_json = "../data/ssq.json"
)

type Bonus struct {
	Name  string `json:"zname"`
	Num   string `json:"znum"`
	Money string `json:"money"`
}

type Info struct {
	Name string `json:"mname"`
}

type Item struct {
	KjIssue string  `json:"kjIssue"`
	KjDate  string  `json:"kjdate"`
	KjZnum  string  `json:"kjznum"`
	KjTnum  string  `json:"kjtnum"`
	MList   []Info  `json:"mlist"`
	Bonus   []Bonus `json:"bonus"`
}

type LotteryContent struct {
	URL       string `json:"url"`
	TextTitle string `json:"textTitle"`
	Lottery   string `json:"lottery"`
	PageNo    string `json:"pageNo"`
	PageSize  string `json:"pageSize"`
	TotalPage string `json:"totalPage"`
	DataList  []Item `json:"dataList"`
}

func TestNewSqlite(t *testing.T) {

	maxCount := 50

	sqlites := make([]*Sqlite3, maxCount)

	var wg sync.WaitGroup
	for i := 0; i < maxCount; i++ {

		wg.Add(1)
		go func(index int) {

			defer wg.Done()

			t.Logf("Open index:%v", index)
			d, err := NewSqlite3(DBPATH)
			if err != nil {
				t.Fatalf("err:%v", err)
			}
			sqlites[index] = d

			array, err := d.GetBlueList("", index)
			if err != nil {
				t.Fatalf("err:%v", err)
			}

			t.Logf("array:%v", array)

		}(i)
	}

	wg.Wait()
	for _, s := range sqlites {
		s.Close()
	}

	t.Logf("success")
}

func TestRefreshHistory(t *testing.T) {
	file, err := os.Open(ssq_json)
	if err != nil {
		t.Fatalf("err:%v", err)
	}

	lottery := LotteryContent{}

	d := json.NewDecoder(file)
	err = d.Decode(&lottery)
	if err != nil {
		t.Fatalf("err:%v", err)
	}

	db, err := NewSqlite3(DBPATH)
	if err != nil {
		t.Fatalf("err:%v", err)
	}

	defer db.Close()

	// remove all history data
	if err := db.RemoveAllHistory(); err != nil {
		t.Fatalf("err:%v", err)
	}

	lotteryList := make([]Lottery, 0)

	for _, item := range lottery.DataList {
		blue, err := strconv.Atoi(item.KjTnum)
		if err != nil {
			t.Fatalf("err:%v", err)
		}

		lotteryList = append(lotteryList, Lottery{
			Expect:   item.KjIssue,
			Red:      item.KjZnum,
			Blue:     blue,
			OpenTime: item.KjDate,
		})
	}

	sort.SliceStable(lotteryList, func(i, j int) bool {
		if strings.Compare(lotteryList[i].Expect, lotteryList[j].Expect) < 0 {
			return true
		}

		return false
	})

	if err := db.BatchAddHistory(lotteryList); err != nil {
		t.Fatalf("err:%v", err)
	}

	t.Logf("success")
}
