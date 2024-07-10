package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2421] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2421][12421] = People_12421
	HistoryPeopleMap[2421][72421] = People_72421
	HistoryPeopleMap[2421][142421] = People_142421
	HistoryPeopleMap[2421][212421] = People_212421
}
