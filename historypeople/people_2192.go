package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2192] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2192][32192] = People_32192
	HistoryPeopleMap[2192][72192] = People_72192
	HistoryPeopleMap[2192][82192] = People_82192
	HistoryPeopleMap[2192][142192] = People_142192
	HistoryPeopleMap[2192][252192] = People_252192
}
