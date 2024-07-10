package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2485] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2485][72485] = People_72485
	HistoryPeopleMap[2485][142485] = People_142485
}
