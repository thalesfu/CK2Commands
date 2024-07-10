package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2499] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2499][72499] = People_72499
	HistoryPeopleMap[2499][142499] = People_142499
	HistoryPeopleMap[2499][212499] = People_212499
}
