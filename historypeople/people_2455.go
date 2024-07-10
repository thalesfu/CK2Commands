package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2455] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2455][72455] = People_72455
	HistoryPeopleMap[2455][142455] = People_142455
}
