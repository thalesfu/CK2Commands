package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2888] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2888][72888] = People_72888
	HistoryPeopleMap[2888][142888] = People_142888
}
