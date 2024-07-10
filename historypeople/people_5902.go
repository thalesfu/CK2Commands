package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5902] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5902][125902] = People_125902
	HistoryPeopleMap[5902][145902] = People_145902
}
