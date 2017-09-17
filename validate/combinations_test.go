package validate

import (
	"reflect"
	"testing"
)

func TestCombinations(t *testing.T) {
	tds := []struct {
		redPrefix  []int
		redPostfix []int
		blue       []int
		expect     []LotteryEntry
	}{
		{ //正确的有前缀+1后缀
			[]int{1, 2, 3, 4, 5, 6},
			[]int{7},
			[]int{8},
			[]LotteryEntry{
				{[6]int{1, 2, 3, 4, 5, 7}, 8},
				{[6]int{1, 2, 3, 4, 6, 7}, 8},
				{[6]int{1, 2, 3, 5, 6, 7}, 8},
				{[6]int{1, 2, 4, 5, 6, 7}, 8},
				{[6]int{1, 3, 4, 5, 6, 7}, 8},
				{[6]int{2, 3, 4, 5, 6, 7}, 8},
			},
		},
		{ //正确的有前缀+0后缀
			[]int{1, 2, 3, 4, 5, 6, 7},
			[]int{},
			[]int{8},
			[]LotteryEntry{
				{[6]int{1, 2, 3, 4, 5, 6}, 8},
				{[6]int{1, 2, 3, 4, 5, 7}, 8},
				{[6]int{1, 2, 3, 4, 6, 7}, 8},
				{[6]int{1, 2, 3, 5, 6, 7}, 8},
				{[6]int{1, 2, 4, 5, 6, 7}, 8},
				{[6]int{1, 3, 4, 5, 6, 7}, 8},
				{[6]int{2, 3, 4, 5, 6, 7}, 8},
			},
		},
		{ //正确的有前缀+0后缀+2蓝球
			[]int{1, 2, 3, 4, 5, 6, 7},
			[]int{},
			[]int{8, 9},
			[]LotteryEntry{
				{[6]int{1, 2, 3, 4, 5, 6}, 8},
				{[6]int{1, 2, 3, 4, 5, 6}, 9},
				{[6]int{1, 2, 3, 4, 5, 7}, 8},
				{[6]int{1, 2, 3, 4, 5, 7}, 9},
				{[6]int{1, 2, 3, 4, 6, 7}, 8},
				{[6]int{1, 2, 3, 4, 6, 7}, 9},
				{[6]int{1, 2, 3, 5, 6, 7}, 8},
				{[6]int{1, 2, 3, 5, 6, 7}, 9},
				{[6]int{1, 2, 4, 5, 6, 7}, 8},
				{[6]int{1, 2, 4, 5, 6, 7}, 9},
				{[6]int{1, 3, 4, 5, 6, 7}, 8},
				{[6]int{1, 3, 4, 5, 6, 7}, 9},
				{[6]int{2, 3, 4, 5, 6, 7}, 8},
				{[6]int{2, 3, 4, 5, 6, 7}, 9},
			},
		},
	}

	for _, td := range tds {
		result := Combinations(td.redPrefix, td.redPostfix, td.blue)
		if !reflect.DeepEqual(result, td.expect) {
			t.Fatalf("recv:%v, expect:%v", result, td.expect)
		}
	}
}
