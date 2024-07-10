package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2799] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2799][72799] = People_72799
	HistoryPeopleMap[2799][142799] = People_142799
	HistoryPeopleMap[2799][222799] = People_222799
}
