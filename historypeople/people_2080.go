package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2080] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2080][12080] = People_12080
	HistoryPeopleMap[2080][72080] = People_72080
	HistoryPeopleMap[2080][82080] = People_82080
	HistoryPeopleMap[2080][142080] = People_142080
	HistoryPeopleMap[2080][252080] = People_252080
	HistoryPeopleMap[2080][412080] = People_412080
}
