package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2381] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2381][12381] = People_12381
	HistoryPeopleMap[2381][72381] = People_72381
	HistoryPeopleMap[2381][142381] = People_142381
	HistoryPeopleMap[2381][212381] = People_212381
	HistoryPeopleMap[2381][252381] = People_252381
	HistoryPeopleMap[2381][462381] = People_462381
}
