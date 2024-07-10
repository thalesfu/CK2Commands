package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[721] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[721][30721] = People_30721
	HistoryPeopleMap[721][70721] = People_70721
	HistoryPeopleMap[721][260721] = People_260721
	HistoryPeopleMap[721][450721] = People_450721
}
