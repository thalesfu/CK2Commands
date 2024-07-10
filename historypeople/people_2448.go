package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2448] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2448][72448] = People_72448
	HistoryPeopleMap[2448][142448] = People_142448
}
