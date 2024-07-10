package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2111] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2111][12111] = People_12111
	HistoryPeopleMap[2111][72111] = People_72111
	HistoryPeopleMap[2111][142111] = People_142111
	HistoryPeopleMap[2111][242111] = People_242111
	HistoryPeopleMap[2111][252111] = People_252111
	HistoryPeopleMap[2111][462111] = People_462111
}
