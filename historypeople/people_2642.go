package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2642] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2642][32642] = People_32642
	HistoryPeopleMap[2642][72642] = People_72642
	HistoryPeopleMap[2642][142642] = People_142642
}
