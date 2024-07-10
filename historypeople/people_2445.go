package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2445] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2445][72445] = People_72445
	HistoryPeopleMap[2445][142445] = People_142445
}
