package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[8259] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[8259][168259] = People_168259
	HistoryPeopleMap[8259][188259] = People_188259
}
