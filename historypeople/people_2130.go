package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2130] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2130][32130] = People_32130
	HistoryPeopleMap[2130][72130] = People_72130
	HistoryPeopleMap[2130][142130] = People_142130
	HistoryPeopleMap[2130][252130] = People_252130
	HistoryPeopleMap[2130][462130] = People_462130
}
