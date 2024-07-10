package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[8604] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[8604][8604] = People_8604
	HistoryPeopleMap[8604][188604] = People_188604
	HistoryPeopleMap[8604][248604] = People_248604
}
