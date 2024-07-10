package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2805] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2805][32805] = People_32805
	HistoryPeopleMap[2805][72805] = People_72805
	HistoryPeopleMap[2805][142805] = People_142805
	HistoryPeopleMap[2805][212805] = People_212805
	HistoryPeopleMap[2805][462805] = People_462805
}
