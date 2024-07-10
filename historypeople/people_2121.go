package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2121] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2121][32121] = People_32121
	HistoryPeopleMap[2121][72121] = People_72121
	HistoryPeopleMap[2121][142121] = People_142121
	HistoryPeopleMap[2121][252121] = People_252121
}
