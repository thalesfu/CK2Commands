package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2377] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2377][12377] = People_12377
	HistoryPeopleMap[2377][72377] = People_72377
	HistoryPeopleMap[2377][142377] = People_142377
	HistoryPeopleMap[2377][252377] = People_252377
	HistoryPeopleMap[2377][462377] = People_462377
}
