package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2483] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2483][72483] = People_72483
	HistoryPeopleMap[2483][142483] = People_142483
}
