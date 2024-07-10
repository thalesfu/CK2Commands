package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2181] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2181][32181] = People_32181
	HistoryPeopleMap[2181][72181] = People_72181
	HistoryPeopleMap[2181][82181] = People_82181
	HistoryPeopleMap[2181][142181] = People_142181
	HistoryPeopleMap[2181][252181] = People_252181
}
