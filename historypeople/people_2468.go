package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2468] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2468][72468] = People_72468
	HistoryPeopleMap[2468][142468] = People_142468
}
