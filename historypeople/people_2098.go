package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2098] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2098][12098] = People_12098
	HistoryPeopleMap[2098][72098] = People_72098
	HistoryPeopleMap[2098][142098] = People_142098
	HistoryPeopleMap[2098][252098] = People_252098
	HistoryPeopleMap[2098][462098] = People_462098
}
