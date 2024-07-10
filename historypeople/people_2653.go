package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2653] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2653][32653] = People_32653
	HistoryPeopleMap[2653][72653] = People_72653
	HistoryPeopleMap[2653][142653] = People_142653
}
