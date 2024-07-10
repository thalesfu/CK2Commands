package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[888] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[888][70888] = People_70888
	HistoryPeopleMap[888][260888] = People_260888
}
