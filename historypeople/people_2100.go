package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2100] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2100][12100] = People_12100
	HistoryPeopleMap[2100][72100] = People_72100
	HistoryPeopleMap[2100][142100] = People_142100
	HistoryPeopleMap[2100][242100] = People_242100
	HistoryPeopleMap[2100][252100] = People_252100
	HistoryPeopleMap[2100][462100] = People_462100
	HistoryPeopleMap[2100][472100] = People_472100
}
