package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[3802] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[3802][73802] = People_73802
	HistoryPeopleMap[3802][93802] = People_93802
}
