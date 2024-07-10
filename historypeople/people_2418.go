package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2418] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2418][12418] = People_12418
	HistoryPeopleMap[2418][72418] = People_72418
	HistoryPeopleMap[2418][142418] = People_142418
}
