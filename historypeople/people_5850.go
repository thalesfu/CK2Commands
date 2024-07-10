package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5850] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5850][145850] = People_145850
	HistoryPeopleMap[5850][205850] = People_205850
	HistoryPeopleMap[5850][455850] = People_455850
}
