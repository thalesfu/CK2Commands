package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2903] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2903][32903] = People_32903
	HistoryPeopleMap[2903][72903] = People_72903
	HistoryPeopleMap[2903][142903] = People_142903
}
