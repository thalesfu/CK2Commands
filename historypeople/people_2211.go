package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2211] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2211][32211] = People_32211
	HistoryPeopleMap[2211][72211] = People_72211
	HistoryPeopleMap[2211][82211] = People_82211
	HistoryPeopleMap[2211][142211] = People_142211
	HistoryPeopleMap[2211][252211] = People_252211
}
