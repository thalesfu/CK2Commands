package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2453] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2453][72453] = People_72453
	HistoryPeopleMap[2453][142453] = People_142453
}
