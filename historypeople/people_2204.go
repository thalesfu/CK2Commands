package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2204] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2204][12204] = People_12204
	HistoryPeopleMap[2204][32204] = People_32204
	HistoryPeopleMap[2204][72204] = People_72204
	HistoryPeopleMap[2204][82204] = People_82204
	HistoryPeopleMap[2204][142204] = People_142204
	HistoryPeopleMap[2204][252204] = People_252204
}
