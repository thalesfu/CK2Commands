package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2389] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2389][12389] = People_12389
	HistoryPeopleMap[2389][72389] = People_72389
	HistoryPeopleMap[2389][142389] = People_142389
	HistoryPeopleMap[2389][252389] = People_252389
}
