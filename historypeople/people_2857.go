package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2857] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2857][72857] = People_72857
	HistoryPeopleMap[2857][142857] = People_142857
}
