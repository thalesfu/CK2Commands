package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[7902] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[7902][7902] = People_7902
	HistoryPeopleMap[7902][167902] = People_167902
	HistoryPeopleMap[7902][247902] = People_247902
}
