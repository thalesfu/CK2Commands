package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5903] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5903][125903] = People_125903
	HistoryPeopleMap[5903][145903] = People_145903
}
