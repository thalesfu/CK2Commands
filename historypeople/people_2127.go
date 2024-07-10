package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2127] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2127][32127] = People_32127
	HistoryPeopleMap[2127][72127] = People_72127
	HistoryPeopleMap[2127][142127] = People_142127
	HistoryPeopleMap[2127][252127] = People_252127
	HistoryPeopleMap[2127][462127] = People_462127
}
