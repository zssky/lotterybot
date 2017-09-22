package main

import (
	"encoding/json"
	"time"

	"github.com/davygeek/lotterybot/db"
	"github.com/juju/errors"
	"github.com/zssky/log"
	httptool "github.com/zssky/tc/http"
)

var (
	host            = "http://f.apiplus.net/ssq-20.json"
	pollDataSeconds = 60 * time.Second
)

type LotteryData struct {
	Expect        string `json:"expect"`
	OpenCode      string `json:"opencode"`
	Opentime      string `json:"opentime"`
	Opentimestamp int64  `json:"opentimestamp"`
}

func main() {

	go Poll()

	dbfile := "./data/db"

	sqlite3, err := db.NewSqlite3(dbfile)
	if err != nil {
		log.Fatalf("%v", err)
	}

	list, err := sqlite3.GetAllHistory(nil, 0)
	if err != nil {
		log.Errorf("%v", errors.ErrorStack(err))
	}

	log.Debugf("list:%v", list)

	blues, err := sqlite3.GetBlueList(0)
	if err != nil {
		log.Errorf("%v", errors.ErrorStack(err))
	}

	log.Debugf("blues:%v", blues)

}

func Poll() {
	pollDataTicker := time.Tick(pollDataSeconds)

	for {
		select {
		case <-pollDataTicker:
			go CollectHistoryData()
		}
	}
}

func CollectHistoryData() {
	data, _, err := httptool.SimpleGet(host, time.Second*10, time.Second*10)
	if err != nil {
		log.Errorf("%v", err)
		return
	}

	var list []LotteryData

	if err := json.Unmarshal(data, &list); err != nil {
		log.Errorf("%v", err)
		return
	}

	log.Debugf("list:%v", list)
}
