package validate

import (
	"strconv"
	"strings"

	"github.com/juju/errors"

	"github.com/zssky/lotterybot/db"
)

type matchKey struct {
	Red  int
	Blue int
}

var (
	moneyMap = map[matchKey]int{
		matchKey{6, 1}: 5000000,
		matchKey{6, 0}: 150000,
		matchKey{5, 1}: 3000,
		matchKey{5, 0}: 200,
		matchKey{4, 1}: 200,
		matchKey{4, 0}: 10,
		matchKey{3, 1}: 10,
		matchKey{2, 1}: 5,
		matchKey{1, 1}: 5,
		matchKey{0, 1}: 5,
	}
)

//ValidateEntryMatch 当前号码命中的结果，RedCount红球命中个数，BlueCount蓝球命中个数.
type ValidateEntryMatch struct {
	RedCount  int
	BlueCount int
	Red       []int
	Blue      int
}

//ValidateEntry 单组中奖结果Entry为选择的号码，Match为命中球的详细信息.
type ValidateEntry struct {
	Entry LotteryEntry
	Match ValidateEntryMatch
	Money int
}

//ValidateResult 验证结果Money是总中奖钱数, History是对应期开奖结果.
type ValidateResult struct {
	Money    int
	Entrys   []ValidateEntry
	History  LotteryEntry
	redPool  [34]bool
	bluePool [17]bool
}

//LotteryEntry 要验证的一组彩票.
type LotteryEntry struct {
	Red  [6]int
	Blue int
}

//Validator 结果验证器.
type Validator struct {
	ldb *db.Sqlite3
}

func NewValidator(dbFile string) (Validator, error) {
	db, err := db.NewSqlite3(dbFile)
	if err != nil {
		return Validator{}, errors.Annotatef(err, "dbFile:%v", dbFile)
	}

	return Validator{ldb: db}, nil
}

//initPool 根据开奖号码初始化兑奖池
func newValidateResult(red string, blue int) (*ValidateResult, error) {
	vr := &ValidateResult{}
	//初始化红球池
	for i, v := range strings.Split(red, ",") {
		r, err := strconv.Atoi(strings.TrimSpace(v))
		if err != nil {
			return vr, errors.Annotatef(err, "invalid red[%d]:%v", i, red)
		}

		if r > 33 || r < 1 {
			return vr, errors.Annotatef(err, "invalid red[%d]:%v, min 1, max 33", i, r)
		}

		vr.History.Red[i] = r
		vr.redPool[r] = true
	}

	//初始化蓝球池
	if blue > 16 || blue < 1 {
		return vr, errors.Errorf("invalid blue:%v, min 1, max 16", blue)
	}

	vr.History.Blue = blue
	vr.bluePool[blue] = true

	return vr, nil
}

//money 根据红球个数及蓝球个数算出来能得多少钱.
func (vr *ValidateResult) money(rc, bc int) (int, error) {
	k := matchKey{Red: rc, Blue: bc}
	v, ok := moneyMap[k]
	if !ok {
		return 0, errors.Errorf("invalid match red count:%v, blue count:%v", rc, bc)
	}

	return v, nil
}

func newValidateEntryMatch(vr *ValidateResult, l LotteryEntry) ValidateEntryMatch {
	vem := ValidateEntryMatch{}

	if vr.bluePool[l.Blue] {
		vem.BlueCount = 1
		vem.Blue = l.Blue
	}

	for _, r := range l.Red {
		if vr.redPool[r] {
			vem.Red = append(vem.Red, r)
			vem.RedCount++
		}
	}

	return vem
}

func (vr *ValidateResult) validate(lotteries []LotteryEntry) error {
	for _, l := range lotteries {
		vem := newValidateEntryMatch(vr, l)
		money, err := vr.money(vem.RedCount, vem.BlueCount)
		if err != nil {
			return errors.Annotatef(err, "lottery:%v", l)
		}
		ve := ValidateEntry{
			Entry: l,
			Match: vem,
			Money: money,
		}

		vr.Money += ve.Money
		vr.Entrys = append(vr.Entrys, ve)
	}

	return nil
}

//Validate 复式结果验证.
func (v Validator) Validate(expect string, lotteries []LotteryEntry) (*ValidateResult, error) {
	ls, err := v.ldb.GetAllHistory(map[string]string{"expect": expect}, 0)
	if err != nil {
		return nil, errors.Trace(err)
	}

	if len(ls) < 1 {
		return nil, errors.Errorf("expect not found:%v", expect)
	}

	vr, err := newValidateResult(ls[0].Red, ls[0].Blue)
	if err != nil {
		return nil, errors.Annotatef(err, "red:%v, blue:%v", ls[0].Red, ls[0].Blue)
	}

	//去计算到底中了多少钱.
	if err = vr.validate(lotteries); err != nil {
		return nil, errors.Trace(err)
	}

	return vr, nil
}
