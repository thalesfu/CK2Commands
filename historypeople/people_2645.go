package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2645] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2645][32645] = People_32645
	HistoryPeopleMap[2645][72645] = People_72645
	HistoryPeopleMap[2645][142645] = People_142645
}
