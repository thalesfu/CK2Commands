package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[903] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[903][903] = People_903
	HistoryPeopleMap[903][40903] = People_40903
	HistoryPeopleMap[903][70903] = People_70903
	HistoryPeopleMap[903][260903] = People_260903
}
