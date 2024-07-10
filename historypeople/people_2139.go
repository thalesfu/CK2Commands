package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2139] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2139][32139] = People_32139
	HistoryPeopleMap[2139][72139] = People_72139
	HistoryPeopleMap[2139][142139] = People_142139
	HistoryPeopleMap[2139][252139] = People_252139
}
