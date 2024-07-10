package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2665] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2665][72665] = People_72665
	HistoryPeopleMap[2665][142665] = People_142665
}
