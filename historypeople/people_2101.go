package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2101] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2101][12101] = People_12101
	HistoryPeopleMap[2101][72101] = People_72101
	HistoryPeopleMap[2101][142101] = People_142101
	HistoryPeopleMap[2101][242101] = People_242101
	HistoryPeopleMap[2101][252101] = People_252101
	HistoryPeopleMap[2101][462101] = People_462101
	HistoryPeopleMap[2101][472101] = People_472101
}
