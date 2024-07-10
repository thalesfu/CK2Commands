package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[7782] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[7782][7782] = People_7782
	HistoryPeopleMap[7782][167782] = People_167782
	HistoryPeopleMap[7782][247782] = People_247782
}
