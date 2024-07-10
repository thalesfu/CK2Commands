package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2608] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2608][32608] = People_32608
	HistoryPeopleMap[2608][72608] = People_72608
	HistoryPeopleMap[2608][142608] = People_142608
	HistoryPeopleMap[2608][452608] = People_452608
}
