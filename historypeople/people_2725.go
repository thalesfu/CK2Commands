package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2725] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2725][32725] = People_32725
	HistoryPeopleMap[2725][72725] = People_72725
	HistoryPeopleMap[2725][142725] = People_142725
}
