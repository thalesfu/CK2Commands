package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2122] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2122][32122] = People_32122
	HistoryPeopleMap[2122][72122] = People_72122
	HistoryPeopleMap[2122][142122] = People_142122
	HistoryPeopleMap[2122][252122] = People_252122
	HistoryPeopleMap[2122][462122] = People_462122
}
