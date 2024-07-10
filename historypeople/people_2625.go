package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2625] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2625][32625] = People_32625
	HistoryPeopleMap[2625][72625] = People_72625
	HistoryPeopleMap[2625][142625] = People_142625
	HistoryPeopleMap[2625][472625] = People_472625
}
