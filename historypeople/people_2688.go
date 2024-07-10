package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2688] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2688][72688] = People_72688
	HistoryPeopleMap[2688][142688] = People_142688
}
