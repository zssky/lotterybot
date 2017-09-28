package db

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/juju/errors"
	//"github.com/zssky/log"
)

type Lottery struct {
	ID            int    `json:"id"`
	Expect        string `json:"expect"`
	Red           string `json:"red"`
	Blue          int    `json:"blue"`
	OpenTime      string `json:"opentime"`
	OpenTimestamp string `json:"opentimestamp"`
}

func (s *Sqlite3) GetAllHistory(query map[string]string, limit int) ([]Lottery, error) {
	sql := "SELECT * FROM history"

	var where string
	for key, value := range query {
		where += fmt.Sprintf("%v=%v AND", key, value)
	}

	if strings.HasSuffix(where, "AND") {
		where = where[:len(where)-len("AND")]
	}

	if where != "" {
		sql += " WHERE " + where
		sql += " ORDER BY expect DESC "
	}

	if limit > 0 {
		sql += fmt.Sprintf("LIMI %v", limit)
	}

	//log.Debugf("sql:%v", sql)
	rows, err := s.Query(sql)
	if err != nil {
		return nil, errors.Trace(err)
	}

	list := make([]Lottery, 0)

	defer rows.Close()
	for rows.Next() {
		var l Lottery

		if err := rows.Scan(&l.ID, &l.Expect, &l.Red, &l.Blue, &l.OpenTime, &l.OpenTimestamp); err != nil {
			return nil, errors.Trace(err)
		}

		list = append(list, l)
	}

	return list, nil
}

func (s *Sqlite3) GetRedList(where string, limit int) ([][]int, error) {
	sql := "SELECT red FROM history"
	if where != "" {
		sql += " WHERE " + where
	}

	if limit > 0 {
		sql += fmt.Sprintf(" ORDER BY expect DESC LIMIT %v", limit)
	}

	//log.Debugf("sql:%v", sql)
	rows, err := s.Query(sql)
	if err != nil {
		return nil, errors.Trace(err)
	}

	list := make([][]int, 0)

	defer rows.Close()

	for rows.Next() {
		var r string
		var l []int

		if err := rows.Scan(&r); err != nil {
			return nil, errors.Trace(err)
		}

		for _, s := range strings.Split(r, ",") {
			num, _ := strconv.Atoi(s)
			l = append(l, num)
		}
		list = append(list, l)
	}

	return list, nil
}

func (s *Sqlite3) GetBlueList(where string, limit int) ([]int, error) {
	sql := "SELECT blue FROM history"
	if where != "" {
		sql += " WHERE " + where
	}

	if limit > 0 {
		sql += fmt.Sprintf(" ORDER BY expect DESC LIMIT %v", limit)
	}

	//log.Debugf("sql:%v", sql)
	rows, err := s.Query(sql)
	if err != nil {
		return nil, errors.Trace(err)
	}

	list := make([]int, 0)

	defer rows.Close()

	for rows.Next() {
		var b int

		if err := rows.Scan(&b); err != nil {
			return nil, errors.Trace(err)
		}
		list = append(list, b)
	}

	return list, nil
}

func (s *Sqlite3) AddHistory(value Lottery) error {
	sql := "INSERT INTO history(expect, red, blue, opentime, opentimestamp) VALUES(?, ?, ?, ?, ?)"

	_, err := s.Exec(sql, value.Expect, value.Red, value.Blue, value.OpenTime, value.OpenTimestamp)
	if err != nil {
		return errors.Trace(err)
	}

	return nil
}

func (s *Sqlite3) BatchAddHistory(list []Lottery) error {
	tx, err := s.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}

	sql := "INSERT INTO history(expect, red, blue, opentime, opentimestamp) VALUES(?, ?, ?, ?, ?)"
	for _, item := range list {
		if _, err := tx.Exec(sql, item.Expect, item.Red, item.Blue, item.OpenTime, item.OpenTimestamp); err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

func (s *Sqlite3) RemoveAllHistory() error {
	sql := "DELETE FROM history"

	if _, err := s.Exec(sql); err != nil {
		return errors.Trace(err)
	}

	sql = "update sqlite_sequence set seq=0 where name='history'"
	if _, err := s.Exec(sql); err != nil {
		return errors.Trace(err)
	}

	return nil
}
