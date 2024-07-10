package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[6700] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[6700][6700] = People_6700
	HistoryPeopleMap[6700][166700] = People_166700
	HistoryPeopleMap[6700][206700] = People_206700
}
