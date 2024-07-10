package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2123] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2123][32123] = People_32123
	HistoryPeopleMap[2123][72123] = People_72123
	HistoryPeopleMap[2123][142123] = People_142123
	HistoryPeopleMap[2123][252123] = People_252123
}
