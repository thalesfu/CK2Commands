package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[599] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[599][599] = People_599
	HistoryPeopleMap[599][30599] = People_30599
	HistoryPeopleMap[599][70599] = People_70599
	HistoryPeopleMap[599][260599] = People_260599
}
