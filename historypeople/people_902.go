package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[902] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[902][902] = People_902
	HistoryPeopleMap[902][30902] = People_30902
	HistoryPeopleMap[902][40902] = People_40902
	HistoryPeopleMap[902][70902] = People_70902
	HistoryPeopleMap[902][260902] = People_260902
}
