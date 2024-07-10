package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[399] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[399][40399] = People_40399
	HistoryPeopleMap[399][190399] = People_190399
	HistoryPeopleMap[399][260399] = People_260399
}
