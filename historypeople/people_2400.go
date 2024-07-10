package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2400] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2400][12400] = People_12400
	HistoryPeopleMap[2400][72400] = People_72400
	HistoryPeopleMap[2400][142400] = People_142400
	HistoryPeopleMap[2400][212400] = People_212400
	HistoryPeopleMap[2400][242400] = People_242400
	HistoryPeopleMap[2400][252400] = People_252400
	HistoryPeopleMap[2400][462400] = People_462400
	HistoryPeopleMap[2400][472400] = People_472400
}
